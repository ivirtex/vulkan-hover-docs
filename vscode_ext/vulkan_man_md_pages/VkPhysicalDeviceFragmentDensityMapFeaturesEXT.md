# VkPhysicalDeviceFragmentDensityMapFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceFragmentDensityMapFeaturesEXT - Structure describing
fragment density map features that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceFragmentDensityMapFeaturesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_fragment_density_map
typedef struct VkPhysicalDeviceFragmentDensityMapFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           fragmentDensityMap;
    VkBool32           fragmentDensityMapDynamic;
    VkBool32           fragmentDensityMapNonSubsampledImages;
} VkPhysicalDeviceFragmentDensityMapFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-fragmentDensityMap"></span> `fragmentDensityMap`
  specifies whether the implementation supports render passes with a
  fragment density map attachment. If this feature is not enabled and
  the `pNext` chain of
  [VkRenderPassCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo.html) includes a
  [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)
  structure, `fragmentDensityMapAttachment` **must** be
  `VK_ATTACHMENT_UNUSED`.

- <span id="features-fragmentDensityMapDynamic"></span>
  `fragmentDensityMapDynamic` specifies whether the implementation
  supports dynamic fragment density map image views. If this feature is
  not enabled,
  `VK_IMAGE_VIEW_CREATE_FRAGMENT_DENSITY_MAP_DYNAMIC_BIT_EXT` **must**
  not be included in
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)::`flags`.

- <span id="features-fragmentDensityMapNonSubsampledImages"></span>
  `fragmentDensityMapNonSubsampledImages` specifies whether the
  implementation supports regular non-subsampled image attachments with
  fragment density map render passes. If this feature is not enabled,
  render passes with a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-fragmentdensitymapattachment"
  target="_blank" rel="noopener">fragment density map attachment</a>
  **must** only have <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-subsamplesampler"
  target="_blank" rel="noopener">subsampled attachments</a> bound.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceFragmentDensityMapFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceFragmentDensityMapFeaturesEXT` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceFragmentDensityMapFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceFragmentDensityMapFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceFragmentDensityMapFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_DENSITY_MAP_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_fragment_density_map](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceFragmentDensityMapFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
