# VkPhysicalDeviceRobustness2PropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceRobustness2PropertiesEXT - Structure describing robust
buffer access properties supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceRobustness2PropertiesEXT` structure is defined as:

``` c
// Provided by VK_EXT_robustness2
typedef struct VkPhysicalDeviceRobustness2PropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       robustStorageBufferAccessSizeAlignment;
    VkDeviceSize       robustUniformBufferAccessSizeAlignment;
} VkPhysicalDeviceRobustness2PropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-robustStorageBufferAccessSizeAlignment"></span>
  `robustStorageBufferAccessSizeAlignment` is the number of bytes that
  the range of a storage buffer descriptor is rounded up to when used
  for bounds-checking when the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a>
  feature is enabled. This value **must** be either 1 or 4.

- <span id="limits-robustUniformBufferAccessSizeAlignment"></span>
  `robustUniformBufferAccessSizeAlignment` is the number of bytes that
  the range of a uniform buffer descriptor is rounded up to when used
  for bounds-checking when the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess2"
  target="_blank" rel="noopener"><code>robustBufferAccess2</code></a>
  feature is enabled. This value **must** be a power of two in the range
  \[1, 256\].

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceRobustness2PropertiesEXT` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceRobustness2PropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceRobustness2PropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceRobustness2PropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ROBUSTNESS_2_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_robustness2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_robustness2.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceRobustness2PropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
