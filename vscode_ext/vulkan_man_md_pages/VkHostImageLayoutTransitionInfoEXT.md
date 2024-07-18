# VkHostImageLayoutTransitionInfoEXT(3) Manual Page

## Name

VkHostImageLayoutTransitionInfoEXT - Structure specifying the parameters
of a host-side image layout transition



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkHostImageLayoutTransitionInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_host_image_copy
typedef struct VkHostImageLayoutTransitionInfoEXT {
    VkStructureType            sType;
    const void*                pNext;
    VkImage                    image;
    VkImageLayout              oldLayout;
    VkImageLayout              newLayout;
    VkImageSubresourceRange    subresourceRange;
} VkHostImageLayoutTransitionInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `image` is a handle to the image affected by this layout transition.

- `oldLayout` is the old layout in an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
  target="_blank" rel="noopener">image layout transition</a>.

- `newLayout` is the new layout in an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-image-layout-transitions"
  target="_blank" rel="noopener">image layout transition</a>.

- `subresourceRange` describes the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views"
  target="_blank" rel="noopener">image subresource range</a> within
  `image` that is affected by this layout transition.

## <a href="#_description" class="anchor"></a>Description

`vkTransitionImageLayoutEXT` does not check whether the device memory
associated with an image is currently in use before performing the
layout transition. The application **must** guarantee that any
previously submitted command that reads from or writes to this
subresource has completed before the host performs the layout
transition. The memory of `image` is accessed by the host as if <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-coherent"
target="_blank" rel="noopener">coherent</a>.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Image layout transitions performed on the host do not require queue
family ownership transfers as the physical layout of the image will not
vary between queue families for the layouts supported by this
function.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>If the device has written to the image memory, it is not
automatically made available to the host. Before this command can be
called, a memory barrier for this image <strong>must</strong> have been
issued on the device with the second <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-dependencies-scopes"
target="_blank" rel="noopener">synchronization scope</a> including
<code>VK_PIPELINE_STAGE_HOST_BIT</code> and
<code>VK_ACCESS_HOST_READ_BIT</code>.</p>
<p>Because queue submissions <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-host-writes"
target="_blank" rel="noopener">automatically make host memory visible to
the device</a>, there would not be a need for a memory barrier before
using the results of this layout transition on the device.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-image-09055"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-image-09055"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-image-09055  
  `image` **must** have been created with
  `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT`

<!-- -->

- <a
  href="#VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01486"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01486"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01486  
  `subresourceRange.baseMipLevel` **must** be less than the `mipLevels`
  specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image`
  was created

- <a
  href="#VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01724"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01724"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01724  
  If `subresourceRange.levelCount` is not `VK_REMAINING_MIP_LEVELS`,
  `subresourceRange.baseMipLevel` + `subresourceRange.levelCount`
  **must** be less than or equal to the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a
  href="#VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01488"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01488"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01488  
  `subresourceRange.baseArrayLayer` **must** be less than the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `image` was created

- <a
  href="#VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01725"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01725"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-01725  
  If `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount`
  **must** be less than or equal to the `arrayLayers` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `image` was created

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-image-01932"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-image-01932"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-image-01932  
  If `image` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-image-09241"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-image-09241"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-image-09241  
  If `image` has a color format that is single-plane, then the
  `aspectMask` member of `subresourceRange` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-image-09242"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-image-09242"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-image-09242  
  If `image` has a color format and is not *disjoint*, then the
  `aspectMask` member of `subresourceRange` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-image-01672"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-image-01672"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-image-01672  
  If `image` has a multi-planar format and the image is *disjoint*, then
  the `aspectMask` member of `subresourceRange` **must** include at
  least one [multi-planar aspect mask](#formats-planes-image-aspect) bit
  or `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-image-03320"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-image-03320"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-image-03320  
  If `image` has a depth/stencil format with both depth and stencil and
  the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is not enabled, then the `aspectMask` member of
  `subresourceRange` **must** include both `VK_IMAGE_ASPECT_DEPTH_BIT`
  and `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-image-03319"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-image-03319"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-image-03319  
  If `image` has a depth/stencil format with both depth and stencil and
  the
  [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts)
  feature is enabled, then the `aspectMask` member of `subresourceRange`
  **must** include either or both `VK_IMAGE_ASPECT_DEPTH_BIT` and
  `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-aspectMask-08702"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-aspectMask-08702"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-aspectMask-08702  
  If the `aspectMask` member of `subresourceRange` includes
  `VK_IMAGE_ASPECT_DEPTH_BIT`, `oldLayout` and `newLayout` **must** not
  be one of `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-aspectMask-08703"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-aspectMask-08703"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-aspectMask-08703  
  If the `aspectMask` member of `subresourceRange` includes
  `VK_IMAGE_ASPECT_STENCIL_BIT`, `oldLayout` and `newLayout` **must**
  not be one of `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`

- <a
  href="#VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-09601"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-09601"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-09601  
  `subresourceRange.aspectMask` **must** be valid for the `format` the
  `image` was created with

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-oldLayout-09229"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-oldLayout-09229"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-oldLayout-09229  
  `oldLayout` **must** be either `VK_IMAGE_LAYOUT_UNDEFINED` or the
  current layout of the image subresources as specified in
  `subresourceRange`

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-oldLayout-09230"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-oldLayout-09230"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-oldLayout-09230  
  If `oldLayout` is not `VK_IMAGE_LAYOUT_UNDEFINED` or
  `VK_IMAGE_LAYOUT_PREINITIALIZED`, it **must** be one of the layouts in
  [VkPhysicalDeviceHostImageCopyPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyPropertiesEXT.html)::`pCopySrcLayouts`

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-newLayout-09057"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-newLayout-09057"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-newLayout-09057  
  `newLayout` **must** be one of the layouts in
  [VkPhysicalDeviceHostImageCopyPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyPropertiesEXT.html)::`pCopyDstLayouts`

Valid Usage (Implicit)

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-sType-sType"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-sType-sType"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_HOST_IMAGE_LAYOUT_TRANSITION_INFO_EXT`

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-pNext-pNext"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-pNext-pNext"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-image-parameter"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-image-parameter"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-oldLayout-parameter"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-oldLayout-parameter"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-oldLayout-parameter  
  `oldLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

- <a href="#VUID-VkHostImageLayoutTransitionInfoEXT-newLayout-parameter"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-newLayout-parameter"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-newLayout-parameter  
  `newLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

- <a
  href="#VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-parameter"
  id="VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-parameter"></a>
  VUID-VkHostImageLayoutTransitionInfoEXT-subresourceRange-parameter  
  `subresourceRange` **must** be a valid
  [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceRange.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkTransitionImageLayoutEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkTransitionImageLayoutEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkHostImageLayoutTransitionInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
