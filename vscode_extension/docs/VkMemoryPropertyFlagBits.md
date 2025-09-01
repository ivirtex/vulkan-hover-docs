# VkMemoryPropertyFlagBits(3) Manual Page

## Name

VkMemoryPropertyFlagBits - Bitmask specifying properties for a memory type



## [](#_c_specification)C Specification

Bits which **may** be set in [VkMemoryType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryType.html)::`propertyFlags`, indicating properties of a memory type, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkMemoryPropertyFlagBits {
    VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT = 0x00000001,
    VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT = 0x00000002,
    VK_MEMORY_PROPERTY_HOST_COHERENT_BIT = 0x00000004,
    VK_MEMORY_PROPERTY_HOST_CACHED_BIT = 0x00000008,
    VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT = 0x00000010,
  // Provided by VK_VERSION_1_1
    VK_MEMORY_PROPERTY_PROTECTED_BIT = 0x00000020,
  // Provided by VK_AMD_device_coherent_memory
    VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD = 0x00000040,
  // Provided by VK_AMD_device_coherent_memory
    VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD = 0x00000080,
  // Provided by VK_NV_external_memory_rdma
    VK_MEMORY_PROPERTY_RDMA_CAPABLE_BIT_NV = 0x00000100,
} VkMemoryPropertyFlagBits;
```

## [](#_description)Description

- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` bit specifies that memory allocated with this type is the most efficient for device access. This property will be set if and only if the memory type belongs to a heap with the `VK_MEMORY_HEAP_DEVICE_LOCAL_BIT` set.
- `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` bit specifies that memory allocated with this type **can** be mapped for host access using [vkMapMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory.html).
- []()`VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` bit specifies that the host cache management commands [vkFlushMappedMemoryRanges](https://registry.khronos.org/vulkan/specs/latest/man/html/vkFlushMappedMemoryRanges.html) and [vkInvalidateMappedMemoryRanges](https://registry.khronos.org/vulkan/specs/latest/man/html/vkInvalidateMappedMemoryRanges.html) are not needed to manage [availability and visibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-available-and-visible) on the host.
- `VK_MEMORY_PROPERTY_HOST_CACHED_BIT` bit specifies that memory allocated with this type is cached on the host. Host memory accesses to uncached memory are slower than to cached memory, however uncached memory is always host coherent.
- `VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT` bit specifies that the memory type only allows device access to the memory. Memory types **must** not have both `VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT` and `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` set. Additionally, the object’s backing memory **may** be provided by the implementation lazily as specified in [Lazily Allocated Memory](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-device-lazy_allocation).
- `VK_MEMORY_PROPERTY_PROTECTED_BIT` bit specifies that the memory type only allows device access to the memory, and allows protected queue operations to access the memory. Memory types **must** not have `VK_MEMORY_PROPERTY_PROTECTED_BIT` set and any of `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` set, or `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` set, or `VK_MEMORY_PROPERTY_HOST_CACHED_BIT` set.
- `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD` bit specifies that device accesses to allocations of this memory type are automatically made [available and visible](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-available-and-visible) on the device. If paired with `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT`, [memory domain operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-available-and-visible) are also performed automatically between host and device.
- `VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD` bit specifies that memory allocated with this type is not cached on the device. Uncached device memory is always device coherent.
- `VK_MEMORY_PROPERTY_RDMA_CAPABLE_BIT_NV` bit specifies that external devices can access this memory directly.

For any memory allocated with both the `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` and the `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD`, host or device accesses also perform automatic memory domain transfer operations, such that writes are always automatically available and visible to both host and device memory domains.

Note

Device coherence is a useful property for certain debugging use cases (e.g. crash analysis, where performing separate coherence actions could mean values are not reported correctly). However, device coherent accesses may be slower than equivalent accesses without device coherence, particularly if they are also device uncached. For device uncached memory in particular, repeated accesses to the same or neighboring memory locations over a short time period (e.g. within a frame) may be slower than it would be for the equivalent cached memory type. As such, it is generally inadvisable to use device coherent or device uncached memory except when really needed.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkMemoryPropertyFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryPropertyFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryPropertyFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0