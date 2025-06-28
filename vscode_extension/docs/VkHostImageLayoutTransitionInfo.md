# VkHostImageLayoutTransitionInfo(3) Manual Page

## Name

VkHostImageLayoutTransitionInfo - Structure specifying the parameters of a host-side image layout transition



## [](#_c_specification)C Specification

The `VkHostImageLayoutTransitionInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkHostImageLayoutTransitionInfo {
    VkStructureType            sType;
    const void*                pNext;
    VkImage                    image;
    VkImageLayout              oldLayout;
    VkImageLayout              newLayout;
    VkImageSubresourceRange    subresourceRange;
} VkHostImageLayoutTransitionInfo;
```

or the equivalent

```c++
// Provided by VK_EXT_host_image_copy
typedef VkHostImageLayoutTransitionInfo VkHostImageLayoutTransitionInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `image` is a handle to the image affected by this layout transition.
- `oldLayout` is the old layout in an [image layout transition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions).
- `newLayout` is the new layout in an [image layout transition](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-image-layout-transitions).
- `subresourceRange` describes the [image subresource range](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views) within `image` that is affected by this layout transition.

## [](#_description)Description

`vkTransitionImageLayout` does not check whether the device memory associated with an image is currently in use before performing the layout transition. The application **must** guarantee that any previously submitted command that reads from or writes to this subresource has completed before the host performs the layout transition. The memory of `image` is accessed by the host as if [coherent](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-coherent).

If `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` and the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature is enabled, the `baseArrayLayer` and `layerCount` members of `subresourceRange` specify the subset of slices of the 3D image affected by the memory barrier, including the layout transition. Any slices of a 3D image not included in `subresourceRange` are not affected by the memory barrier and remain in their existing layout.

Note

Image layout transitions performed on the host do not require queue family ownership transfers as the physical layout of the image will not vary between queue families for the layouts supported by this function.

Note

If the device has written to the image memory, it is not automatically made available to the host. Before this command can be called, a memory barrier for this image **must** have been issued on the device with the second [synchronization scope](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies-scopes) including `VK_PIPELINE_STAGE_HOST_BIT` and `VK_ACCESS_HOST_READ_BIT`.

Because queue submissions [automatically make host memory visible to the device](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-submission-host-writes), there would not be a need for a memory barrier before using the results of this layout transition on the device.

Valid Usage

- [](#VUID-VkHostImageLayoutTransitionInfo-image-09055)VUID-VkHostImageLayoutTransitionInfo-image-09055  
  `image` **must** have been created with `VK_IMAGE_USAGE_HOST_TRANSFER_BIT`

<!--THE END-->

- [](#VUID-VkHostImageLayoutTransitionInfo-subresourceRange-01486)VUID-VkHostImageLayoutTransitionInfo-subresourceRange-01486  
  `subresourceRange.baseMipLevel` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkHostImageLayoutTransitionInfo-subresourceRange-01724)VUID-VkHostImageLayoutTransitionInfo-subresourceRange-01724  
  If `subresourceRange.levelCount` is not `VK_REMAINING_MIP_LEVELS`, `subresourceRange.baseMipLevel` + `subresourceRange.levelCount` **must** be less than or equal to the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkHostImageLayoutTransitionInfo-subresourceRange-01488)VUID-VkHostImageLayoutTransitionInfo-subresourceRange-01488  
  If `image` is not a 3D image or was created without `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, or the [`maintenance9`](#features-maintenance9) feature is not enabled, `subresourceRange.baseArrayLayer` **must** be less than the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkHostImageLayoutTransitionInfo-maintenance9-10798)VUID-VkHostImageLayoutTransitionInfo-maintenance9-10798  
  If the [`maintenance9`](#features-maintenance9) feature is enabled and `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, `subresourceRange.baseArrayLayer` **must** be less than the depth computed from `baseMipLevel` and `extent.depth` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created, according to the formula defined in [Image Mip Level Sizing](#resources-image-mip-level-sizing)
- [](#VUID-VkHostImageLayoutTransitionInfo-maintenance9-10799)VUID-VkHostImageLayoutTransitionInfo-maintenance9-10799  
  If the [`maintenance9`](#features-maintenance9) feature is enabled and `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, `subresourceRange.levelCount` **must** be 1
- [](#VUID-VkHostImageLayoutTransitionInfo-subresourceRange-01725)VUID-VkHostImageLayoutTransitionInfo-subresourceRange-01725  
  If `image` is not a 3D image or was created without `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, or the [`maintenance9`](#features-maintenance9) feature is not enabled, and `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created
- [](#VUID-VkHostImageLayoutTransitionInfo-maintenance9-10800)VUID-VkHostImageLayoutTransitionInfo-maintenance9-10800  
  If the [`maintenance9`](#features-maintenance9) feature is enabled, `subresourceRange.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, and `image` is a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT` set, `subresourceRange.baseArrayLayer` + `subresourceRange.layerCount` **must** be less than or equal to the depth computed from `baseMipLevel` and `extent.depth` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `image` was created, according to the formula defined in [Image Mip Level Sizing](#resources-image-mip-level-sizing)
- [](#VUID-VkHostImageLayoutTransitionInfo-image-01932)VUID-VkHostImageLayoutTransitionInfo-image-01932  
  If `image` is non-sparse then it **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkHostImageLayoutTransitionInfo-image-09241)VUID-VkHostImageLayoutTransitionInfo-image-09241  
  If `image` has a color format that is single-plane, then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkHostImageLayoutTransitionInfo-image-09242)VUID-VkHostImageLayoutTransitionInfo-image-09242  
  If `image` has a color format and is not *disjoint*, then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkHostImageLayoutTransitionInfo-image-01672)VUID-VkHostImageLayoutTransitionInfo-image-01672  
  If `image` has a [multi-planar format](#formats-multiplanar) and the image is *disjoint*, then the `aspectMask` member of `subresourceRange` **must** include at least one [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit or `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkHostImageLayoutTransitionInfo-image-03320)VUID-VkHostImageLayoutTransitionInfo-image-03320  
  If `image` has a depth/stencil format with both depth and stencil and the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is not enabled, then the `aspectMask` member of `subresourceRange` **must** include both `VK_IMAGE_ASPECT_DEPTH_BIT` and `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkHostImageLayoutTransitionInfo-image-03319)VUID-VkHostImageLayoutTransitionInfo-image-03319  
  If `image` has a depth/stencil format with both depth and stencil and the [`separateDepthStencilLayouts`](#features-separateDepthStencilLayouts) feature is enabled, then the `aspectMask` member of `subresourceRange` **must** include either or both `VK_IMAGE_ASPECT_DEPTH_BIT` and `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkHostImageLayoutTransitionInfo-image-10749)VUID-VkHostImageLayoutTransitionInfo-image-10749  
  If `image` has a depth-only format then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_DEPTH_BIT`
- [](#VUID-VkHostImageLayoutTransitionInfo-image-10750)VUID-VkHostImageLayoutTransitionInfo-image-10750  
  If `image` has a stencil-only format then the `aspectMask` member of `subresourceRange` **must** be `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkHostImageLayoutTransitionInfo-aspectMask-08702)VUID-VkHostImageLayoutTransitionInfo-aspectMask-08702  
  If the `aspectMask` member of `subresourceRange` includes `VK_IMAGE_ASPECT_DEPTH_BIT`, `oldLayout` and `newLayout` **must** not be one of `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkHostImageLayoutTransitionInfo-aspectMask-08703)VUID-VkHostImageLayoutTransitionInfo-aspectMask-08703  
  If the `aspectMask` member of `subresourceRange` includes `VK_IMAGE_ASPECT_STENCIL_BIT`, `oldLayout` and `newLayout` **must** not be one of `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL`
- [](#VUID-VkHostImageLayoutTransitionInfo-subresourceRange-09601)VUID-VkHostImageLayoutTransitionInfo-subresourceRange-09601  
  `subresourceRange.aspectMask` **must** be valid for the `format` the `image` was created with
- [](#VUID-VkHostImageLayoutTransitionInfo-oldLayout-09229)VUID-VkHostImageLayoutTransitionInfo-oldLayout-09229  
  `oldLayout` **must** be either `VK_IMAGE_LAYOUT_UNDEFINED` or the current layout of the image subresources as specified in `subresourceRange`
- [](#VUID-VkHostImageLayoutTransitionInfo-oldLayout-09230)VUID-VkHostImageLayoutTransitionInfo-oldLayout-09230  
  If `oldLayout` is not `VK_IMAGE_LAYOUT_UNDEFINED`, `VK_IMAGE_LAYOUT_ZERO_INITIALIZED_EXT`, or `VK_IMAGE_LAYOUT_PREINITIALIZED`, it **must** be one of the layouts in [VkPhysicalDeviceHostImageCopyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceHostImageCopyProperties.html)::`pCopySrcLayouts`
- [](#VUID-VkHostImageLayoutTransitionInfo-newLayout-09057)VUID-VkHostImageLayoutTransitionInfo-newLayout-09057  
  `newLayout` **must** be one of the layouts in [VkPhysicalDeviceHostImageCopyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceHostImageCopyProperties.html)::`pCopyDstLayouts`

Valid Usage (Implicit)

- [](#VUID-VkHostImageLayoutTransitionInfo-sType-sType)VUID-VkHostImageLayoutTransitionInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_HOST_IMAGE_LAYOUT_TRANSITION_INFO`
- [](#VUID-VkHostImageLayoutTransitionInfo-pNext-pNext)VUID-VkHostImageLayoutTransitionInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkHostImageLayoutTransitionInfo-image-parameter)VUID-VkHostImageLayoutTransitionInfo-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkHostImageLayoutTransitionInfo-oldLayout-parameter)VUID-VkHostImageLayoutTransitionInfo-oldLayout-parameter  
  `oldLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkHostImageLayoutTransitionInfo-newLayout-parameter)VUID-VkHostImageLayoutTransitionInfo-newLayout-parameter  
  `newLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkHostImageLayoutTransitionInfo-subresourceRange-parameter)VUID-VkHostImageLayoutTransitionInfo-subresourceRange-parameter  
  `subresourceRange` **must** be a valid [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html) structure

## [](#_see_also)See Also

[VK\_EXT\_host\_image\_copy](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_host_image_copy.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkImageSubresourceRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSubresourceRange.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkTransitionImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/vkTransitionImageLayout.html), [vkTransitionImageLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkTransitionImageLayoutEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkHostImageLayoutTransitionInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0