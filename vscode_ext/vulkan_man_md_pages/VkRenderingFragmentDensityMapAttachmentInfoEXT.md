# VkRenderingFragmentDensityMapAttachmentInfoEXT(3) Manual Page

## Name

VkRenderingFragmentDensityMapAttachmentInfoEXT - Structure specifying
fragment shading rate attachment information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderingFragmentDensityMapAttachmentInfoEXT` structure is
defined as:

``` c
// Provided by VK_KHR_dynamic_rendering with VK_EXT_fragment_density_map
typedef struct VkRenderingFragmentDensityMapAttachmentInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkImageView        imageView;
    VkImageLayout      imageLayout;
} VkRenderingFragmentDensityMapAttachmentInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `imageView` is the image view that will be used as a fragment density
  map attachment.

- `imageLayout` is the layout that `imageView` will be in during
  rendering.

## <a href="#_description" class="anchor"></a>Description

This structure can be included in the `pNext` chain of
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) to define a fragment density
map. If this structure is not included in the `pNext` chain, `imageView`
is treated as [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html).

Valid Usage

- <a
  href="#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06157"
  id="VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06157"></a>
  VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06157  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageLayout` **must** be `VK_IMAGE_LAYOUT_GENERAL` or
  `VK_IMAGE_LAYOUT_FRAGMENT_DENSITY_MAP_OPTIMAL_EXT`

- <a
  href="#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06158"
  id="VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06158"></a>
  VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06158  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it
  **must** have been created with
  `VK_IMAGE_USAGE_FRAGMENT_DENSITY_MAP_BIT_EXT`

- <a
  href="#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06159"
  id="VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06159"></a>
  VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-06159  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it
  **must** not have been created with
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

- <a
  href="#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-apiVersion-07908"
  id="VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-apiVersion-07908"></a>
  VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-apiVersion-07908  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview"
  target="_blank" rel="noopener"><code>multiview</code></a> feature is
  not enabled,
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, and `imageView` is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it **must** have a `layerCount`
  equal to `1`

Valid Usage (Implicit)

- <a
  href="#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-sType-sType"
  id="VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-sType-sType"></a>
  VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDERING_FRAGMENT_DENSITY_MAP_ATTACHMENT_INFO_EXT`

- <a
  href="#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-parameter"
  id="VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-parameter"></a>
  VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageView-parameter  
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

- <a
  href="#VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageLayout-parameter"
  id="VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageLayout-parameter"></a>
  VUID-VkRenderingFragmentDensityMapAttachmentInfoEXT-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_fragment_density_map](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_fragment_density_map.html),
[VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderingFragmentDensityMapAttachmentInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
