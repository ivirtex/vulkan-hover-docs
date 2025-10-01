# VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR(3) Manual Page

## Name

VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR - Structure describing supported queues for indirect copy



## [](#_c_specification)C Specification

The `VkPhysicalDeviceCopyMemoryIndirectPropertiesNV` structure is defined as:

```c++
// Provided by VK_KHR_copy_memory_indirect
typedef struct VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkQueueFlags       supportedQueues;
} VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR;
```

or the equivalent

```c++
// Provided by VK_NV_copy_memory_indirect
typedef VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR VkPhysicalDeviceCopyMemoryIndirectPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `supportedQueues` is a bitmask of [VkQueueFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFlagBits.html) indicating the types of queues on which [indirect copy commands](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#indirect-copies) are supported. If a queue family supports any of the bits set in `supportedQueues`, then it **must** support at least one [indirect copy command](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#indirect-copies).

## [](#_description)Description

If the [`indirectMemoryCopy`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-indirectMemoryCopy) or [`indirectMemoryToImageCopy`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-indirectMemoryToImageCopy) feature is supported, `supportedQueues` **must** return at least one supported queue type.

If the `VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR-sType-sType)VUID-VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COPY_MEMORY_INDIRECT_PROPERTIES_KHR`

## [](#_see_also)See Also

[VK\_KHR\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_memory_indirect.html), [VkQueueFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceCopyMemoryIndirectPropertiesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0