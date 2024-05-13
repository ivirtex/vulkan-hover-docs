# VkPhysicalDeviceFragmentDensityMapPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceFragmentDensityMapPropertiesEXT - Structure describing
fragment density map properties that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFragmentDensityMapPropertiesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_fragment_density_map
typedef struct VkPhysicalDeviceFragmentDensityMapPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         minFragmentDensityTexelSize;
    VkExtent2D         maxFragmentDensityTexelSize;
    VkBool32           fragmentDensityInvocations;
} VkPhysicalDeviceFragmentDensityMapPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-minFragmentDensityTexelSize"></span>
  `minFragmentDensityTexelSize` is the minimum <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-fragment-density-texel-size"
  target="_blank" rel="noopener">fragment density texel size</a>.

- <span id="limits-maxFragmentDensityTexelSize"></span>
  `maxFragmentDensityTexelSize` is the maximum fragment density texel
  size.

- <span id="limits-fragmentDensityInvocations"></span>
  `fragmentDensityInvocations` specifies whether the implementation
  **may** invoke additional fragment shader invocations for each covered
  sample.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceFragmentDensityMapPropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceFragmentDensityMapPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceFragmentDensityMapPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceFragmentDensityMapPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_fragment_density_map](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFragmentDensityMapPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
