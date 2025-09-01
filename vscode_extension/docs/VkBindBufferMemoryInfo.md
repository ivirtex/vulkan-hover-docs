# VkBindBufferMemoryInfo(3) Manual Page

## Name

VkBindBufferMemoryInfo - Structure specifying how to bind a buffer to memory



## [](#_c_specification)C Specification

`VkBindBufferMemoryInfo` contains members corresponding to the parameters of [vkBindBufferMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindBufferMemory.html).

The `VkBindBufferMemoryInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkBindBufferMemoryInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkBuffer           buffer;
    VkDeviceMemory     memory;
    VkDeviceSize       memoryOffset;
} VkBindBufferMemoryInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_bind_memory2
typedef VkBindBufferMemoryInfo VkBindBufferMemoryInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `buffer` is the buffer to be attached to memory.
- `memory` is a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) object describing the device memory to attach.
- `memoryOffset` is the start offset of the region of `memory` which is to be bound to the buffer. The number of bytes returned in the `VkMemoryRequirements`::`size` member in `memory`, starting from `memoryOffset` bytes, will be bound to the specified buffer.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBindBufferMemoryInfo-buffer-07459)VUID-VkBindBufferMemoryInfo-buffer-07459  
  `buffer` **must** not have been bound to a memory object
- [](#VUID-VkBindBufferMemoryInfo-buffer-01030)VUID-VkBindBufferMemoryInfo-buffer-01030  
  `buffer` **must** not have been created with any sparse memory binding flags
- [](#VUID-VkBindBufferMemoryInfo-memoryOffset-01031)VUID-VkBindBufferMemoryInfo-memoryOffset-01031  
  `memoryOffset` **must** be less than the size of `memory`
- [](#VUID-VkBindBufferMemoryInfo-memory-01035)VUID-VkBindBufferMemoryInfo-memory-01035  
  `memory` **must** have been allocated using one of the memory types allowed in the `memoryTypeBits` member of the `VkMemoryRequirements` structure returned from a call to `vkGetBufferMemoryRequirements` with `buffer`
- [](#VUID-VkBindBufferMemoryInfo-None-10739)VUID-VkBindBufferMemoryInfo-None-10739  
  If `memory` was not allocated from a memory heap with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property set, `memoryOffset` **must** be an integer multiple of the `alignment` member of the `VkMemoryRequirements` structure returned from a call to `vkGetBufferMemoryRequirements` with `buffer`
- [](#VUID-VkBindBufferMemoryInfo-memory-10740)VUID-VkBindBufferMemoryInfo-memory-10740  
  If `memory` was allocated from a memory heap with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property set, `memoryOffset` **must** be an integer multiple of the `alignment` member of the `VkTileMemoryRequirementsQCOM` structure returned from a call to `vkGetBufferMemoryRequirements` with `buffer`
- [](#VUID-VkBindBufferMemoryInfo-None-10741)VUID-VkBindBufferMemoryInfo-None-10741  
  If `memory` was not allocated from a memory heap with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property set, `size` member of the `VkMemoryRequirements` structure returned from a call to `vkGetBufferMemoryRequirements` with `buffer` **must** be less than or equal to the size of `memory` minus `memoryOffset`
- [](#VUID-VkBindBufferMemoryInfo-memory-10742)VUID-VkBindBufferMemoryInfo-memory-10742  
  If `memory` was allocated from a memory heap with the `VK_MEMORY_HEAP_TILE_MEMORY_BIT_QCOM` property set, `size` member of the `VkTileMemoryRequirementsQCOM` structure returned from a call to `vkGetBufferMemoryRequirements` with `buffer` **must** be less than or equal to the size of `memory` minus `memoryOffset`
- [](#VUID-VkBindBufferMemoryInfo-buffer-01444)VUID-VkBindBufferMemoryInfo-buffer-01444  
  If `buffer` requires a dedicated allocation (as reported by [vkGetBufferMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferMemoryRequirements2.html) in [VkMemoryDedicatedRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedRequirements.html)::`requiresDedicatedAllocation` for `buffer`), `memory` **must** have been allocated with [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer` equal to `buffer`
- [](#VUID-VkBindBufferMemoryInfo-memory-01508)VUID-VkBindBufferMemoryInfo-memory-01508  
  If the `VkMemoryAllocateInfo` provided when `memory` was allocated included a [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfo.html) structure in its `pNext` chain, and [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer` was not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), then `buffer` **must** equal [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer`, and `memoryOffset` **must** be zero
- [](#VUID-VkBindBufferMemoryInfo-None-01898)VUID-VkBindBufferMemoryInfo-None-01898  
  If `buffer` was created with the `VK_BUFFER_CREATE_PROTECTED_BIT` bit set, the buffer **must** be bound to a memory object allocated with a memory type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`
- [](#VUID-VkBindBufferMemoryInfo-None-01899)VUID-VkBindBufferMemoryInfo-None-01899  
  If `buffer` was created with the `VK_BUFFER_CREATE_PROTECTED_BIT` bit not set, the buffer **must** not be bound to a memory object allocated with a memory type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`
- [](#VUID-VkBindBufferMemoryInfo-buffer-01038)VUID-VkBindBufferMemoryInfo-buffer-01038  
  If `buffer` was created with [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDedicatedAllocationBufferCreateInfoNV.html)::`dedicatedAllocation` equal to `VK_TRUE`, `memory` **must** have been allocated with [VkDedicatedAllocationMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDedicatedAllocationMemoryAllocateInfoNV.html)::`buffer` equal to a buffer handle created with identical creation parameters to `buffer` and `memoryOffset` **must** be zero
- [](#VUID-VkBindBufferMemoryInfo-apiVersion-07920)VUID-VkBindBufferMemoryInfo-apiVersion-07920  
  If the [VK\_KHR\_dedicated\_allocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dedicated_allocation.html) extension is not enabled, [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, and `buffer` was not created with [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDedicatedAllocationBufferCreateInfoNV.html)::`dedicatedAllocation` equal to `VK_TRUE`, `memory` **must** not have been allocated dedicated for a specific buffer or image
- [](#VUID-VkBindBufferMemoryInfo-memory-02726)VUID-VkBindBufferMemoryInfo-memory-02726  
  If the value of [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes` used to allocate `memory` is not `0`, it **must** include at least one of the handles set in [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryBufferCreateInfo.html)::`handleTypes` when `buffer` was created
- [](#VUID-VkBindBufferMemoryInfo-memory-02985)VUID-VkBindBufferMemoryInfo-memory-02985  
  If `memory` was allocated by a memory import operation, that is not [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportAndroidHardwareBufferInfoANDROID.html) with a non-`NULL` `buffer` value, the external handle type of the imported memory **must** also have been set in [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryBufferCreateInfo.html)::`handleTypes` when `buffer` was created
- [](#VUID-VkBindBufferMemoryInfo-memory-02986)VUID-VkBindBufferMemoryInfo-memory-02986  
  If `memory` was allocated with the [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportAndroidHardwareBufferInfoANDROID.html) memory import operation with a non-`NULL` `buffer` value, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID` **must** also have been set in [VkExternalMemoryBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryBufferCreateInfo.html)::`handleTypes` when `buffer` was created
- [](#VUID-VkBindBufferMemoryInfo-bufferDeviceAddress-03339)VUID-VkBindBufferMemoryInfo-bufferDeviceAddress-03339  
  If the [VkPhysicalDeviceBufferDeviceAddressFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBufferDeviceAddressFeatures.html)::`bufferDeviceAddress` feature is enabled and `buffer` was created with the `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT` bit set, `memory` **must** have been allocated with the `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT` bit set
- [](#VUID-VkBindBufferMemoryInfo-bufferDeviceAddressCaptureReplay-09200)VUID-VkBindBufferMemoryInfo-bufferDeviceAddressCaptureReplay-09200  
  If the [VkPhysicalDeviceBufferDeviceAddressFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceBufferDeviceAddressFeatures.html)::`bufferDeviceAddressCaptureReplay` feature is enabled and `buffer` was created with the `VK_BUFFER_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` bit set, `memory` **must** have been allocated with the `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` bit set
- [](#VUID-VkBindBufferMemoryInfo-buffer-06408)VUID-VkBindBufferMemoryInfo-buffer-06408  
  If `buffer` was created with [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html) chained to [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`pNext`, `memory` **must** be allocated with a [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryBufferCollectionFUCHSIA.html) chained to [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html)::`pNext`
- [](#VUID-VkBindBufferMemoryInfo-descriptorBufferCaptureReplay-08112)VUID-VkBindBufferMemoryInfo-descriptorBufferCaptureReplay-08112  
  If the `buffer` was created with the `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` bit set, `memory` **must** have been allocated with the `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT` bit set
- [](#VUID-VkBindBufferMemoryInfo-buffer-09201)VUID-VkBindBufferMemoryInfo-buffer-09201  
  If the `buffer` was created with the `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` bit set, `memory` **must** have been allocated with the `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` bit set
- [](#VUID-VkBindBufferMemoryInfo-pNext-01605)VUID-VkBindBufferMemoryInfo-pNext-01605  
  If the `pNext` chain includes a [VkBindBufferMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryDeviceGroupInfo.html) structure, all instances of `memory` specified by [VkBindBufferMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryDeviceGroupInfo.html)::`pDeviceIndices` **must** have been allocated

Valid Usage (Implicit)

- [](#VUID-VkBindBufferMemoryInfo-sType-sType)VUID-VkBindBufferMemoryInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_BUFFER_MEMORY_INFO`
- [](#VUID-VkBindBufferMemoryInfo-pNext-pNext)VUID-VkBindBufferMemoryInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkBindBufferMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindBufferMemoryDeviceGroupInfo.html) or [VkBindMemoryStatus](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindMemoryStatus.html)
- [](#VUID-VkBindBufferMemoryInfo-sType-unique)VUID-VkBindBufferMemoryInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkBindBufferMemoryInfo-buffer-parameter)VUID-VkBindBufferMemoryInfo-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkBindBufferMemoryInfo-memory-parameter)VUID-VkBindBufferMemoryInfo-memory-parameter  
  `memory` **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html) handle
- [](#VUID-VkBindBufferMemoryInfo-commonparent)VUID-VkBindBufferMemoryInfo-commonparent  
  Both of `buffer`, and `memory` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `buffer` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkBindBufferMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindBufferMemory2.html), [vkBindBufferMemory2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindBufferMemory2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindBufferMemoryInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0