# vkAllocateMemory(3) Manual Page

## Name

vkAllocateMemory - Allocate device memory



## [](#_c_specification)C Specification

To allocate memory objects, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkAllocateMemory(
    VkDevice                                    device,
    const VkMemoryAllocateInfo*                 pAllocateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDeviceMemory*                             pMemory);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the memory.
- `pAllocateInfo` is a pointer to a [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure describing parameters of the allocation. A successfully returned allocation **must** use the requested parameters — no substitution is permitted by the implementation.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pMemory` is a pointer to a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle in which information about the allocated memory is returned.

## [](#_description)Description

Allocations returned by `vkAllocateMemory` are guaranteed to meet any alignment requirement of the implementation. For example, if an implementation requires 128 byte alignment for images and 64 byte alignment for buffers, the device memory returned through this mechanism would be 128-byte aligned. This ensures that applications **can** correctly suballocate objects of different types (with potentially different alignment requirements) in the same memory object.

When memory is allocated, its contents are undefined with the following constraint:

- The contents of unprotected memory **must** not be a function of the contents of data protected memory objects, even if those memory objects were previously freed.

Note

The contents of memory allocated by one application **should** not be a function of data from protected memory objects of another application, even if those memory objects were previously freed.

The maximum number of valid memory allocations that **can** exist simultaneously within a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) **may** be restricted by implementation- or platform-dependent limits. The [`maxMemoryAllocationCount`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxMemoryAllocationCount) feature describes the number of allocations that **can** exist simultaneously before encountering these internal limits.

Note

For historical reasons, if `maxMemoryAllocationCount` is exceeded, some implementations may return `VK_ERROR_TOO_MANY_OBJECTS`. Exceeding this limit will result in undefined behavior, and an application should not rely on the use of the returned error code in order to identify when the limit is reached.

Note

Many protected memory implementations involve complex hardware and system software support, and often have additional and much lower limits on the number of simultaneous protected memory allocations (from memory types with the `VK_MEMORY_PROPERTY_PROTECTED_BIT` property) than for non-protected memory allocations. These limits can be system-wide, and depend on a variety of factors outside of the Vulkan implementation, so they cannot be queried in Vulkan. Applications **should** use as few allocations as possible from such memory types by suballocating aggressively, and be prepared for allocation failure even when there is apparently plenty of capacity remaining in the memory heap. As a guideline, the Vulkan conformance test suite requires that at least 80 minimum-size allocations can exist concurrently when no other uses of protected memory are active in the system.

Some platforms **may** have a limit on the maximum size of a single allocation. For example, certain systems **may** fail to create allocations with a size greater than or equal to 4GB. Such a limit is implementation-dependent, and if such a failure occurs then the error `VK_ERROR_OUT_OF_DEVICE_MEMORY` **must** be returned. This limit is advertised in [VkPhysicalDeviceMaintenance3Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance3Properties.html)::`maxMemoryAllocationSize`.

The cumulative memory size allocated to a heap **can** be limited by the size of the specified heap. In such cases, allocated memory is tracked on a per-device and per-heap basis. Some platforms allow overallocation into other heaps. The overallocation behavior **can** be specified through the `VK_AMD_memory_overallocation_behavior` extension.

If the `memoryTypeIndex` belongs to a heap with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` bit included in its properties, this allocation is backed by tile memory, which is an on device cache. Unlike other heaps, allocations out of the tile memory will always have a starting address at the start of the heap and its contents are aliased with all other [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) objects bound to the same range while executing within the same [*tile memory scope*](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-tile-heaps).

If the [VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePageableDeviceLocalMemoryFeaturesEXT.html)::`pageableDeviceLocalMemory` feature is enabled, memory allocations made from a heap that includes `VK_MEMORY_HEAP_DEVICE_LOCAL_BIT` in [VkMemoryHeap](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeap.html)::`flags` **may** be transparently moved to host-local memory allowing multiple applications to share device-local memory. If there is no space left in device-local memory when this new allocation is made, other allocations **may** be moved out transparently to make room. The operating system will determine which allocations to move to device-local memory or host-local memory based on platform-specific criteria. To help the operating system make good choices, the application **should** set the appropriate memory priority with [VkMemoryPriorityAllocateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryPriorityAllocateInfoEXT.html) and adjust it as necessary with [vkSetDeviceMemoryPriorityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetDeviceMemoryPriorityEXT.html). Higher priority allocations will moved to device-local memory first.

Memory allocations made on heaps without the `VK_MEMORY_HEAP_DEVICE_LOCAL_BIT` property will not be transparently promoted to device-local memory by the operating system.

Valid Usage

- [](#VUID-vkAllocateMemory-pAllocateInfo-01713)VUID-vkAllocateMemory-pAllocateInfo-01713  
  `pAllocateInfo->allocationSize` **must** be less than or equal to [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html)::`memoryHeaps`\[memindex].`size` where `memindex` = [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html)::`memoryTypes`\[`pAllocateInfo->memoryTypeIndex`].`heapIndex` as returned by [vkGetPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceMemoryProperties.html) for the [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) that `device` was created from
- [](#VUID-vkAllocateMemory-pAllocateInfo-01714)VUID-vkAllocateMemory-pAllocateInfo-01714  
  `pAllocateInfo->memoryTypeIndex` **must** be less than [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMemoryProperties.html)::`memoryTypeCount` as returned by [vkGetPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceMemoryProperties.html) for the [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) that `device` was created from
- [](#VUID-vkAllocateMemory-deviceCoherentMemory-02790)VUID-vkAllocateMemory-deviceCoherentMemory-02790  
  If the [`deviceCoherentMemory`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceCoherentMemory) feature is not enabled, `pAllocateInfo->memoryTypeIndex` **must** not identify a memory type supporting `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD`
- [](#VUID-vkAllocateMemory-maxMemoryAllocationCount-04101)VUID-vkAllocateMemory-maxMemoryAllocationCount-04101  
  There **must** be less than `VkPhysicalDeviceLimits`::`maxMemoryAllocationCount` device memory allocations currently allocated on the device
- [](#VUID-VkTileMemoryBindInfoQCOM-memoryTypeIndex)VUID-VkTileMemoryBindInfoQCOM-memoryTypeIndex  
  If the [`tileMemoryHeap`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tileMemoryHeap) feature is not enabled, `pAllocateInfo->memoryTypeIndex` **must** not identify a memory type that corresponds to a [VkMemoryHeap](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHeap.html) with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property

Valid Usage (Implicit)

- [](#VUID-vkAllocateMemory-device-parameter)VUID-vkAllocateMemory-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkAllocateMemory-pAllocateInfo-parameter)VUID-vkAllocateMemory-pAllocateInfo-parameter  
  `pAllocateInfo` **must** be a valid pointer to a valid [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure
- [](#VUID-vkAllocateMemory-pAllocator-parameter)VUID-vkAllocateMemory-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkAllocateMemory-pMemory-parameter)VUID-vkAllocateMemory-pMemory-parameter  
  `pMemory` **must** be a valid pointer to a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-vkAllocateMemory-device-queuecount)VUID-vkAllocateMemory-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_EXTERNAL_HANDLE`
- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkAllocateMemory)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0