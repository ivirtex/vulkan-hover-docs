# VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV - Structure describing sparse address space limits of an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_extended_sparse_address_space
typedef struct VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV {
    VkStructureType       sType;
    void*                 pNext;
    VkDeviceSize          extendedSparseAddressSpaceSize;
    VkImageUsageFlags     extendedSparseImageUsageFlags;
    VkBufferUsageFlags    extendedSparseBufferUsageFlags;
} VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`extendedSparseAddressSpaceSize` is the total amount of address space available, in bytes, for sparse memory resources of all usages if the [`extendedSparseAddressSpace`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-extendedSparseAddressSpace) feature is enabled. This **must** be greater than or equal to `VkPhysicalDeviceLimits`::`sparseAddressSpaceSize`, and the difference in space **must** only be used with usages allowed below. This is an upper bound on the sum of the sizes of all sparse resources, regardless of whether any memory is bound to them.
- []()`extendedSparseImageUsageFlags` is a bitmask of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) of usages which **may** allow an implementation to use the full `extendedSparseAddressSpaceSize` space.
- []()`extendedSparseBufferUsageFlags` is a bitmask of [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlagBits.html) of usages which **may** allow an implementation to use the full `extendedSparseAddressSpaceSize` space.

## [](#_description)Description

If the `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV-sType-sType)VUID-VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_SPARSE_ADDRESS_SPACE_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_extended\_sparse\_address\_space](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_extended_sparse_address_space.html), [VkBufferUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferUsageFlags.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0