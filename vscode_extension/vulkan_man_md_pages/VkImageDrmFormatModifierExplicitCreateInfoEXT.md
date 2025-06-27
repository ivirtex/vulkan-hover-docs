# VkImageDrmFormatModifierExplicitCreateInfoEXT(3) Manual Page

## Name

VkImageDrmFormatModifierExplicitCreateInfoEXT - Specify that an image be
created with the provided DRM format modifier and explicit memory layout



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
includes a
[VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)
structure, then the image will be created with the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-drm-format-modifier"
target="_blank" rel="noopener">Linux DRM format modifier</a> and memory
layout defined by the structure.

The
[VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html)
structure is defined as:

``` c
// Provided by VK_EXT_image_drm_format_modifier
typedef struct VkImageDrmFormatModifierExplicitCreateInfoEXT {
    VkStructureType               sType;
    const void*                   pNext;
    uint64_t                      drmFormatModifier;
    uint32_t                      drmFormatModifierPlaneCount;
    const VkSubresourceLayout*    pPlaneLayouts;
} VkImageDrmFormatModifierExplicitCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `drmFormatModifier` is the *Linux DRM format modifier* with which the
  image will be created.

- `drmFormatModifierPlaneCount` is the number of *memory planes* in the
  image (as reported by
  [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html))
  as well as the length of the `pPlaneLayouts` array.

- `pPlaneLayouts` is a pointer to an array of
  [VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html) structures describing
  the image’s *memory planes*.

## <a href="#_description" class="anchor"></a>Description

The `i`<sup>th</sup> member of `pPlaneLayouts` describes the layout of
the image’s `i`<sup>th</sup> *memory plane* (that is,
`VK_IMAGE_ASPECT_MEMORY_PLANE`*`_i_`*`BIT_EXT`). In each element of
`pPlaneLayouts`, the implementation **must** ignore `size`. The
implementation calculates the size of each plane, which the application
**can** query with
[vkGetImageSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageSubresourceLayout.html).

When creating an image with
[VkImageDrmFormatModifierExplicitCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageDrmFormatModifierExplicitCreateInfoEXT.html),
it is the application’s responsibility to satisfy all valid usage
requirements. However, the implementation **must** validate that the
provided `pPlaneLayouts`, when combined with the provided
`drmFormatModifier` and other creation parameters in
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) and its `pNext` chain,
produce a valid image. (This validation is necessarily
implementation-dependent and outside the scope of Vulkan, and therefore
not described by valid usage requirements). If this validation fails,
then [vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html) returns
`VK_ERROR_INVALID_DRM_FORMAT_MODIFIER_PLANE_LAYOUT_EXT`.

Valid Usage

- <a
  href="#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifier-02264"
  id="VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifier-02264"></a>
  VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifier-02264  
  `drmFormatModifier` **must** be compatible with the parameters in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) and its `pNext` chain, as
  determined by querying
  [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageFormatInfo2.html)
  extended with
  [VkPhysicalDeviceImageDrmFormatModifierInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageDrmFormatModifierInfoEXT.html)

- <a
  href="#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifierPlaneCount-02265"
  id="VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifierPlaneCount-02265"></a>
  VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifierPlaneCount-02265  
  `drmFormatModifierPlaneCount` **must** be equal to the
  [VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)::`drmFormatModifierPlaneCount`
  associated with [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`format`
  and `drmFormatModifier`, as found by querying
  [VkDrmFormatModifierPropertiesListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesListEXT.html)

- <a href="#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-size-02267"
  id="VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-size-02267"></a>
  VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-size-02267  
  For each element of `pPlaneLayouts`, `size` **must** be 0

- <a
  href="#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-arrayPitch-02268"
  id="VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-arrayPitch-02268"></a>
  VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-arrayPitch-02268  
  For each element of `pPlaneLayouts`, `arrayPitch` **must** be 0 if
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`arrayLayers` is 1

- <a
  href="#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-depthPitch-02269"
  id="VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-depthPitch-02269"></a>
  VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-depthPitch-02269  
  For each element of `pPlaneLayouts`, `depthPitch` **must** be 0 if
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`extent.depth` is 1

Valid Usage (Implicit)

- <a
  href="#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-sType-sType"
  id="VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-sType-sType"></a>
  VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_DRM_FORMAT_MODIFIER_EXPLICIT_CREATE_INFO_EXT`

- <a
  href="#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-pPlaneLayouts-parameter"
  id="VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-pPlaneLayouts-parameter"></a>
  VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-pPlaneLayouts-parameter  
  `pPlaneLayouts` **must** be a valid pointer to an array of
  `drmFormatModifierPlaneCount`
  [VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html) structures

- <a
  href="#VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifierPlaneCount-arraylength"
  id="VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifierPlaneCount-arraylength"></a>
  VUID-VkImageDrmFormatModifierExplicitCreateInfoEXT-drmFormatModifierPlaneCount-arraylength  
  `drmFormatModifierPlaneCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_drm_format_modifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_drm_format_modifier.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubresourceLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubresourceLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageDrmFormatModifierExplicitCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
