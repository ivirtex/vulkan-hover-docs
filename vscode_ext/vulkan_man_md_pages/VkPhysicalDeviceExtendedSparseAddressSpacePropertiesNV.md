# VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV - Structure
describing sparse address space limits of an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV` structure
is defined as:

``` c
// Provided by VK_NV_extended_sparse_address_space
typedef struct VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV {
    VkStructureType       sType;
    void*                 pNext;
    VkDeviceSize          extendedSparseAddressSpaceSize;
    VkImageUsageFlags     extendedSparseImageUsageFlags;
    VkBufferUsageFlags    extendedSparseBufferUsageFlags;
} VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-extendedSparseAddressSpaceSize"></span>
  `extendedSparseAddressSpaceSize` is the total amount of address space
  available, in bytes, for sparse memory resources of all usages if the
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-extendedSparseAddressSpace"
  target="_blank"
  rel="noopener"><code>extendedSparseAddressSpace</code></a> feature is
  enabled. This **must** be greater than or equal to
  `VkPhysicalDeviceLimits`::`sparseAddressSpaceSize`, and the difference
  in space **must** only be used with usages allowed below. This is an
  upper bound on the sum of the sizes of all sparse resources,
  regardless of whether any memory is bound to them.

- <span id="limits-extendedSparseImageUsageFlags"></span>
  `extendedSparseImageUsageFlags` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) of usages which
  **may** allow an implementation to use the full
  `extendedSparseAddressSpaceSize` space.

- <span id="limits-extendedSparseBufferUsageFlags"></span>
  `extendedSparseBufferUsageFlags` is a bitmask of
  [VkBufferUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlagBits.html) of usages which
  **may** allow an implementation to use the full
  `extendedSparseAddressSpaceSize` space.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV`
structure is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_SPARSE_ADDRESS_SPACE_PROPERTIES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_extended_sparse_address_space](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_extended_sparse_address_space.html),
[VkBufferUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
