# VkImageSubresourceRange(3) Manual Page

## Name

VkImageSubresourceRange - Structure specifying an image subresource range



## [](#_c_specification)C Specification

The `VkImageSubresourceRange` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkImageSubresourceRange {
    VkImageAspectFlags    aspectMask;
    uint32_t              baseMipLevel;
    uint32_t              levelCount;
    uint32_t              baseArrayLayer;
    uint32_t              layerCount;
} VkImageSubresourceRange;
```

## [](#_members)Members

- `aspectMask` is a bitmask of [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) specifying which aspect(s) of the image are included in the view.
- `baseMipLevel` is the first mipmap level accessible to the view.
- `levelCount` is the number of mipmap levels (starting from `baseMipLevel`) accessible to the view.
- `baseArrayLayer` is the first array layer accessible to the view.
- `layerCount` is the number of array layers (starting from `baseArrayLayer`) accessible to the view.

## [](#_description)Description

The number of mipmap levels and array layers **must** be a subset of the image subresources in the image. If an application wants to use all mip levels or layers in an image after the `baseMipLevel` or `baseArrayLayer`, it **can** set `levelCount` and `layerCount` to the special values `VK_REMAINING_MIP_LEVELS` and `VK_REMAINING_ARRAY_LAYERS` without knowing the exact number of mip levels or layers.

For cube and cube array image views, the layers of the image view starting at `baseArrayLayer` correspond to faces in the order +X, -X, +Y, -Y, +Z, -Z. For cube arrays, each set of six sequential layers is a single cube, so the number of cube maps in a cube map array view is *`layerCount` / 6*, and image array layer (`baseArrayLayer` + i) is face index (i mod 6) of cube *i / 6*. If the number of layers in the view, whether set explicitly in `layerCount` or implied by `VK_REMAINING_ARRAY_LAYERS`, is not a multiple of 6, the last cube map in the array **must** not be accessed.

`aspectMask` **must** be only `VK_IMAGE_ASPECT_COLOR_BIT`, `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT` if `format` is a color, depth-only or stencil-only format, respectively, except if `format` is a [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar). If using a depth/stencil format with both depth and stencil components, `aspectMask` **must** include at least one of `VK_IMAGE_ASPECT_DEPTH_BIT` and `VK_IMAGE_ASPECT_STENCIL_BIT`, and **can** include both.

When the `VkImageSubresourceRange` structure is used to select a subset of the slices of a 3D image’s mip level in order to create a 2D or 2D array image view of a 3D image created with `VK_IMAGE_CREATE_2D_ARRAY_COMPATIBLE_BIT`, `baseArrayLayer` and `layerCount` specify the first slice index and the number of slices to include in the created image view. Such an image view **can** be used as a framebuffer attachment that refers only to the specified range of slices of the selected mip level. If the [`maintenance9`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance9) feature is not enabled, any layout transitions performed on such an attachment view during a render pass instance still apply to the entire subresource referenced which includes all the slices of the selected mip level.

When using an image view of a depth/stencil image to populate a descriptor set (e.g. for sampling in the shader, or for use as an input attachment), the `aspectMask` **must** only include one bit, which selects whether the image view is used for depth reads (i.e. using a floating-point sampler or input attachment in the shader) or stencil reads (i.e. using an unsigned integer sampler or input attachment in the shader). When an image view of a depth/stencil image is used as a depth/stencil framebuffer attachment, the `aspectMask` is ignored and both depth and stencil image subresources are used.

When creating a `VkImageView`, if [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) is enabled in the sampler, the `aspectMask` of a `subresourceRange` used by the `VkImageView` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`.

When creating a `VkImageView`, if sampler Y′CBCR conversion is not enabled in the sampler and the image `format` is [multi-planar format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar), the image **must** have been created with `VK_IMAGE_CREATE_MUTABLE_FORMAT_BIT`, and the `aspectMask` of the `VkImageView`’s `subresourceRange` **must** be `VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT` or `VK_IMAGE_ASPECT_PLANE_2_BIT`.

Valid Usage

- [](#VUID-VkImageSubresourceRange-levelCount-01720)VUID-VkImageSubresourceRange-levelCount-01720  
  If `levelCount` is not `VK_REMAINING_MIP_LEVELS`, it **must** be greater than `0`
- [](#VUID-VkImageSubresourceRange-layerCount-01721)VUID-VkImageSubresourceRange-layerCount-01721  
  If `layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, it **must** be greater than `0`
- [](#VUID-VkImageSubresourceRange-aspectMask-01670)VUID-VkImageSubresourceRange-aspectMask-01670  
  If `aspectMask` includes `VK_IMAGE_ASPECT_COLOR_BIT`, then it **must** not include any of `VK_IMAGE_ASPECT_PLANE_0_BIT`, `VK_IMAGE_ASPECT_PLANE_1_BIT`, or `VK_IMAGE_ASPECT_PLANE_2_BIT`
- [](#VUID-VkImageSubresourceRange-aspectMask-02278)VUID-VkImageSubresourceRange-aspectMask-02278  
  `aspectMask` **must** not include `VK_IMAGE_ASPECT_MEMORY_PLANE_i_BIT_EXT` for any index *i*

Valid Usage (Implicit)

- [](#VUID-VkImageSubresourceRange-aspectMask-parameter)VUID-VkImageSubresourceRange-aspectMask-parameter  
  `aspectMask` **must** be a valid combination of [VkImageAspectFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlagBits.html) values
- [](#VUID-VkImageSubresourceRange-aspectMask-requiredbitmask)VUID-VkImageSubresourceRange-aspectMask-requiredbitmask  
  `aspectMask` **must** not be `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkHostImageLayoutTransitionInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkHostImageLayoutTransitionInfo.html), [VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageAspectFlags.html), [VkImageMemoryBarrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier.html), [VkImageMemoryBarrier2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageMemoryBarrier2.html), [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewCreateInfo.html), [vkCmdClearColorImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdClearColorImage.html), [vkCmdClearDepthStencilImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdClearDepthStencilImage.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageSubresourceRange)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0