# vkCmdCopyImage(3) Manual Page

## Name

vkCmdCopyImage - Copy data between images



## [](#_c_specification)C Specification

To copy data between image objects, call:

```c++
// Provided by VK_VERSION_1_0
void vkCmdCopyImage(
    VkCommandBuffer                             commandBuffer,
    VkImage                                     srcImage,
    VkImageLayout                               srcImageLayout,
    VkImage                                     dstImage,
    VkImageLayout                               dstImageLayout,
    uint32_t                                    regionCount,
    const VkImageCopy*                          pRegions);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `srcImage` is the source image.
- `srcImageLayout` is the current layout of the source image subresource.
- `dstImage` is the destination image.
- `dstImageLayout` is the current layout of the destination image subresource.
- `regionCount` is the number of regions to copy.
- `pRegions` is a pointer to an array of [VkImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy.html) structures specifying the regions to copy.

## [](#_description)Description

Each source region specified by `pRegions` is copied from the source image to the destination region of the destination image. If any of the specified regions in `srcImage` overlaps in memory with any of the specified regions in `dstImage`, values read from those overlapping regions are undefined. If any region accesses a depth aspect in `dstImage` and the `VK_EXT_depth_range_unrestricted` extension is not enabled, values copied from `srcBuffer` outside of the range \[0,1] will be written as undefined values to the destination image.

[Multi-planar images](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar) **can** only be copied on a per-plane basis, and the subresources used in each region when copying to or from such images **must** specify only one plane, though different regions **can** specify different planes. When copying planes of multi-planar images, the format considered is the [compatible format for that plane](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatible-planes), rather than the format of the multi-planar image.

If the format of the destination image has a different [block extent](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-compatibility-classes) than the source image (e.g. one is a compressed format), the offset and extent for each of the regions specified is [scaled according to the block extents of each format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-size-compatibility) to match in size. Copy regions for each image **must** be aligned to a multiple of the texel block extent in each dimension, except at the edges of the image, where region extents **must** match the edge of the image.

Image data **can** be copied between images with different image types. If one image is `VK_IMAGE_TYPE_3D` and the other image is `VK_IMAGE_TYPE_2D` with multiple layers, then each slice is copied to or from a different layer; `depth` slices in the 3D image correspond to `layerCount` layers in the 2D image, with an effective `depth` of `1` used for the 2D image. If the [`maintenance5`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-maintenance5) feature is enabled, all other combinations are allowed and function as if 1D images are 2D images with a height of 1. Otherwise, other combinations of image types are disallowed.

Valid Usage

- [](#VUID-vkCmdCopyImage-commandBuffer-01825)VUID-vkCmdCopyImage-commandBuffer-01825  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `srcImage` **must** not be a protected image
- [](#VUID-vkCmdCopyImage-commandBuffer-01826)VUID-vkCmdCopyImage-commandBuffer-01826  
  If `commandBuffer` is an unprotected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstImage` **must** not be a protected image
- [](#VUID-vkCmdCopyImage-commandBuffer-01827)VUID-vkCmdCopyImage-commandBuffer-01827  
  If `commandBuffer` is a protected command buffer and [`protectedNoFault`](#limits-protectedNoFault) is not supported, `dstImage` **must** not be an unprotected image
- [](#VUID-vkCmdCopyImage-commandBuffer-10217)VUID-vkCmdCopyImage-commandBuffer-10217  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each element of `pRegions`, where the `aspectMask` member of `srcSubresource` is `VK_IMAGE_ASPECT_COLOR_BIT`, the `aspectMask` of `dstSubresource` **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-vkCmdCopyImage-commandBuffer-10218)VUID-vkCmdCopyImage-commandBuffer-10218  
  If the queue family used to create the [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html) which `commandBuffer` was allocated from does not support `VK_QUEUE_GRAPHICS_BIT`, for each element of `pRegions`, where the `aspectMask` member of `dstSubresource` is `VK_IMAGE_ASPECT_COLOR_BIT` then the `aspectMask` of `srcSubresource` **must** not be `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

<!--THE END-->

- [](#VUID-vkCmdCopyImage-pRegions-00124)VUID-vkCmdCopyImage-pRegions-00124  
  The union of all source regions, and the union of all destination regions, specified by the elements of `pRegions`, **must** not overlap in memory
- [](#VUID-vkCmdCopyImage-srcImage-01995)VUID-vkCmdCopyImage-srcImage-01995  
  The [format features](#resources-image-format-features) of `srcImage` **must** contain `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT`
- [](#VUID-vkCmdCopyImage-srcImageLayout-00128)VUID-vkCmdCopyImage-srcImageLayout-00128  
  `srcImageLayout` **must** specify the layout of the image subresources of `srcImage` specified in `pRegions` at the time this command is executed on a `VkDevice`
- [](#VUID-vkCmdCopyImage-srcImageLayout-01917)VUID-vkCmdCopyImage-srcImageLayout-01917  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-vkCmdCopyImage-srcImage-09460)VUID-vkCmdCopyImage-srcImage-09460  
  If `srcImage` and `dstImage` are the same, and any elements of `pRegions` contains the `srcSubresource` and `dstSubresource` with matching `mipLevel` and overlapping array layers, then the `srcImageLayout` and `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`
- [](#VUID-vkCmdCopyImage-dstImage-01996)VUID-vkCmdCopyImage-dstImage-01996  
  The [format features](#resources-image-format-features) of `dstImage` **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`
- [](#VUID-vkCmdCopyImage-dstImageLayout-00133)VUID-vkCmdCopyImage-dstImageLayout-00133  
  `dstImageLayout` **must** specify the layout of the image subresources of `dstImage` specified in `pRegions` at the time this command is executed on a `VkDevice`
- [](#VUID-vkCmdCopyImage-dstImageLayout-01395)VUID-vkCmdCopyImage-dstImageLayout-01395  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-vkCmdCopyImage-srcImage-01548)VUID-vkCmdCopyImage-srcImage-01548  
  If the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each of `srcImage` and `dstImage` is not a [multi-planar format](#formats-multiplanar), the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each of `srcImage` and `dstImage` **must** be [size-compatible](#formats-size-compatibility)
- [](#VUID-vkCmdCopyImage-None-01549)VUID-vkCmdCopyImage-None-01549  
  In a copy to or from a plane of a [multi-planar image](#formats-multiplanar), the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of the image and plane **must** be compatible according to [the description of compatible planes](#formats-compatible-planes) for the plane being copied
- [](#VUID-vkCmdCopyImage-srcImage-09247)VUID-vkCmdCopyImage-srcImage-09247  
  If the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each of `srcImage` and `dstImage` is a [compressed image format](#compressed_image_formats), the formats **must** have the same texel block extent
- [](#VUID-vkCmdCopyImage-srcImage-00136)VUID-vkCmdCopyImage-srcImage-00136  
  The sample count of `srcImage` and `dstImage` **must** match
- [](#VUID-vkCmdCopyImage-srcOffset-01783)VUID-vkCmdCopyImage-srcOffset-01783  
  The `srcOffset` and `extent` members of each element of `pRegions` **must** respect the image transfer granularity requirements of `commandBuffer`’s command pool’s queue family, as described in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html)
- [](#VUID-vkCmdCopyImage-dstOffset-01784)VUID-vkCmdCopyImage-dstOffset-01784  
  The `dstOffset` and `extent` members of each element of `pRegions` **must** respect the image transfer granularity requirements of `commandBuffer`’s command pool’s queue family, as described in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html)
- [](#VUID-vkCmdCopyImage-srcImage-01551)VUID-vkCmdCopyImage-srcImage-01551  
  If neither `srcImage` nor `dstImage` has a [multi-planar format](#formats-multiplanar) and the [`maintenance8`](#features-maintenance8) feature is not enabled then for each element of `pRegions`, `srcSubresource.aspectMask` and `dstSubresource.aspectMask` **must** match
- [](#VUID-vkCmdCopyImage-srcSubresource-10214)VUID-vkCmdCopyImage-srcSubresource-10214  
  If `srcSubresource.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT`, then `dstSubresource.aspectMask` **must** not contain both `VK_IMAGE_ASPECT_DEPTH_BIT` and `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-vkCmdCopyImage-dstSubresource-10215)VUID-vkCmdCopyImage-dstSubresource-10215  
  If `dstSubresource.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT`, then `srcSubresource.aspectMask` **must** not contain both `VK_IMAGE_ASPECT_DEPTH_BIT` and `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-vkCmdCopyImage-srcImage-08713)VUID-vkCmdCopyImage-srcImage-08713  
  If `srcImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `srcSubresource.aspectMask` **must** be a single valid [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit
- [](#VUID-vkCmdCopyImage-dstImage-08714)VUID-vkCmdCopyImage-dstImage-08714  
  If `dstImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `dstSubresource.aspectMask` **must** be a single valid [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit
- [](#VUID-vkCmdCopyImage-srcImage-01556)VUID-vkCmdCopyImage-srcImage-01556  
  If `srcImage` has a [multi-planar format](#formats-multiplanar) and the `dstImage` does not have a multi-planar image format, then for each element of `pRegions`, `dstSubresource.aspectMask` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-vkCmdCopyImage-dstImage-01557)VUID-vkCmdCopyImage-dstImage-01557  
  If `dstImage` has a [multi-planar format](#formats-multiplanar) and the `srcImage` does not have a multi-planar image format, then for each element of `pRegions`, `srcSubresource.aspectMask` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-vkCmdCopyImage-srcSubresource-10211)VUID-vkCmdCopyImage-srcSubresource-10211  
  If `srcSubresource.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT` and `dstSubresource.aspectMask` is `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`, then the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values of `srcImage` and `dstImage` **must** be compatible according to [the list of compatible depth-stencil and color formats](#formats-compatible-zs-color)
- [](#VUID-vkCmdCopyImage-srcSubresource-10212)VUID-vkCmdCopyImage-srcSubresource-10212  
  If `srcSubresource.aspectMask` is `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT` and `dstSubresource.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT`, then the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values of `srcImage` and `dstImage` **must** be compatible according to [the list of compatible depth-stencil and color formats](#formats-compatible-zs-color)
- [](#VUID-vkCmdCopyImage-apiVersion-07932)VUID-vkCmdCopyImage-apiVersion-07932  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled, or [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, and either `srcImage` or `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `srcSubresource.baseArrayLayer` and `dstSubresource.baseArrayLayer` **must** both be `0`, and `srcSubresource.layerCount` and `dstSubresource.layerCount` **must** both be `1`
- [](#VUID-vkCmdCopyImage-srcImage-04443)VUID-vkCmdCopyImage-srcImage-04443  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `srcSubresource.baseArrayLayer` **must** be `0` and `srcSubresource.layerCount` **must** be `1`
- [](#VUID-vkCmdCopyImage-dstImage-04444)VUID-vkCmdCopyImage-dstImage-04444  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `dstSubresource.baseArrayLayer` **must** be `0` and `dstSubresource.layerCount` **must** be `1`
- [](#VUID-vkCmdCopyImage-aspectMask-00142)VUID-vkCmdCopyImage-aspectMask-00142  
  For each element of `pRegions`, `srcSubresource.aspectMask` **must** specify aspects present in `srcImage`
- [](#VUID-vkCmdCopyImage-aspectMask-00143)VUID-vkCmdCopyImage-aspectMask-00143  
  For each element of `pRegions`, `dstSubresource.aspectMask` **must** specify aspects present in `dstImage`
- [](#VUID-vkCmdCopyImage-srcOffset-00144)VUID-vkCmdCopyImage-srcOffset-00144  
  For each element of `pRegions`, `srcOffset.x` and (`extent.width` + `srcOffset.x`) **must** both be greater than or equal to `0` and less than or equal to the width of the specified `srcSubresource` of `srcImage`
- [](#VUID-vkCmdCopyImage-srcOffset-00145)VUID-vkCmdCopyImage-srcOffset-00145  
  For each element of `pRegions`, `srcOffset.y` and (`extent.height` + `srcOffset.y`) **must** both be greater than or equal to `0` and less than or equal to the height of the specified `srcSubresource` of `srcImage`
- [](#VUID-vkCmdCopyImage-srcImage-00146)VUID-vkCmdCopyImage-srcImage-00146  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `srcOffset.y` **must** be `0` and `extent.height` **must** be `1`
- [](#VUID-vkCmdCopyImage-srcOffset-00147)VUID-vkCmdCopyImage-srcOffset-00147  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `srcOffset.z` and (`extent.depth` + `srcOffset.z`) **must** both be greater than or equal to `0` and less than or equal to the depth of the specified `srcSubresource` of `srcImage`
- [](#VUID-vkCmdCopyImage-srcImage-01785)VUID-vkCmdCopyImage-srcImage-01785  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `srcOffset.z` **must** be `0` and `extent.depth` **must** be `1`
- [](#VUID-vkCmdCopyImage-dstImage-01786)VUID-vkCmdCopyImage-dstImage-01786  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `dstOffset.z` **must** be `0`
- [](#VUID-vkCmdCopyImage-srcImage-10907)VUID-vkCmdCopyImage-srcImage-10907  
  If either the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each of `srcImage` and `dstImage` is not a [compressed image format](#compressed_image_formats), and `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `extent.depth` **must** be `1`
- [](#VUID-vkCmdCopyImage-srcImage-01787)VUID-vkCmdCopyImage-srcImage-01787  
  If `srcImage` is of type `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `srcOffset.z` **must** be `0`
- [](#VUID-vkCmdCopyImage-dstImage-01788)VUID-vkCmdCopyImage-dstImage-01788  
  If `dstImage` is of type `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `dstOffset.z` **must** be `0`
- [](#VUID-vkCmdCopyImage-apiVersion-07933)VUID-vkCmdCopyImage-apiVersion-07933  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled, and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, `srcImage` and `dstImage` **must** have the same [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html)
- [](#VUID-vkCmdCopyImage-apiVersion-08969)VUID-vkCmdCopyImage-apiVersion-08969  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled, and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, `srcImage` or `dstImage` is of type `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `extent.depth` **must** be `1`
- [](#VUID-vkCmdCopyImage-srcImage-07743)VUID-vkCmdCopyImage-srcImage-07743  
  If `srcImage` and `dstImage` have a different [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html), and the [`maintenance5`](#features-maintenance5) feature is not enabled, one **must** be `VK_IMAGE_TYPE_3D` and the other **must** be `VK_IMAGE_TYPE_2D`
- [](#VUID-vkCmdCopyImage-srcImage-08793)VUID-vkCmdCopyImage-srcImage-08793  
  If `srcImage` and `dstImage` have the same [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html), for each element of `pRegions`, if neither of the `layerCount` members of `srcSubresource` or `dstSubresource` are `VK_REMAINING_ARRAY_LAYERS`, the `layerCount` members of `srcSubresource` or `dstSubresource` **must** match
- [](#VUID-vkCmdCopyImage-srcImage-08794)VUID-vkCmdCopyImage-srcImage-08794  
  If `srcImage` and `dstImage` have the same [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html), and one of the `layerCount` members of `srcSubresource` or `dstSubresource` is `VK_REMAINING_ARRAY_LAYERS`, the other member **must** be either `VK_REMAINING_ARRAY_LAYERS` or equal to the `arrayLayers` member of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) used to create the image minus `baseArrayLayer`
- [](#VUID-vkCmdCopyImage-srcImage-01790)VUID-vkCmdCopyImage-srcImage-01790  
  If `srcImage` and `dstImage` are both of type `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `extent.depth` **must** be `1`
- [](#VUID-vkCmdCopyImage-srcImage-01791)VUID-vkCmdCopyImage-srcImage-01791  
  If `srcImage` is of type `VK_IMAGE_TYPE_2D`, and `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `extent.depth` **must** equal `srcSubresource.layerCount`
- [](#VUID-vkCmdCopyImage-dstImage-01792)VUID-vkCmdCopyImage-dstImage-01792  
  If `dstImage` is of type `VK_IMAGE_TYPE_2D`, and `srcImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `extent.depth` **must** equal `dstSubresource.layerCount`
- [](#VUID-vkCmdCopyImage-dstOffset-00150)VUID-vkCmdCopyImage-dstOffset-00150  
  For each element of `pRegions`, `dstOffset.x` and (`extent.width` + `dstOffset.x`), where `extent` is [adjusted for size-compatibility](#formats-size-compatibility), **must** both be greater than or equal to `0` and less than or equal to the width of the specified `dstSubresource` of `dstImage`
- [](#VUID-vkCmdCopyImage-dstOffset-00151)VUID-vkCmdCopyImage-dstOffset-00151  
  For each element of `pRegions`, `dstOffset.y` and (`extent.height` + `dstOffset.y`), where `extent` is [adjusted for size-compatibility](#formats-size-compatibility), **must** both be greater than or equal to `0` and less than or equal to the height of the specified `dstSubresource` of `dstImage`
- [](#VUID-vkCmdCopyImage-dstImage-00152)VUID-vkCmdCopyImage-dstImage-00152  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `dstOffset.y` **must** be `0`
- [](#VUID-vkCmdCopyImage-srcImage-10908)VUID-vkCmdCopyImage-srcImage-10908  
  If either the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each of `srcImage` and `dstImage` is not a [compressed image format](#compressed_image_formats), and `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `extent.height` **must** be `1`, where `extent` is [adjusted for size-compatibility](#formats-size-compatibility)
- [](#VUID-vkCmdCopyImage-dstOffset-00153)VUID-vkCmdCopyImage-dstOffset-00153  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `dstOffset.z` and (`extent.depth` + `dstOffset.z`), where `extent` is [adjusted for size-compatibility](#formats-size-compatibility), **must** both be greater than or equal to `0` and less than or equal to the depth of the specified `dstSubresource` of `dstImage`
- [](#VUID-vkCmdCopyImage-pRegions-07278)VUID-vkCmdCopyImage-pRegions-07278  
  For each element of `pRegions`, `srcOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImage-pRegions-07279)VUID-vkCmdCopyImage-pRegions-07279  
  For each element of `pRegions`, `srcOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImage-pRegions-07280)VUID-vkCmdCopyImage-pRegions-07280  
  For each element of `pRegions`, `srcOffset.z` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImage-pRegions-07281)VUID-vkCmdCopyImage-pRegions-07281  
  For each element of `pRegions`, `dstOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyImage-pRegions-07282)VUID-vkCmdCopyImage-pRegions-07282  
  For each element of `pRegions`, `dstOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyImage-pRegions-07283)VUID-vkCmdCopyImage-pRegions-07283  
  For each element of `pRegions`, `dstOffset.z` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-vkCmdCopyImage-srcImage-01728)VUID-vkCmdCopyImage-srcImage-01728  
  For each element of `pRegions`, if the sum of `srcOffset.x` and `extent.width` does not equal the width of the subresource specified by `srcSubresource`, `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImage-srcImage-01729)VUID-vkCmdCopyImage-srcImage-01729  
  For each element of `pRegions`, if the sum of `srcOffset.y` and `extent.height` does not equal the height of the subresource specified by `srcSubresource`, `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImage-srcImage-01730)VUID-vkCmdCopyImage-srcImage-01730  
  For each element of `pRegions`, if the sum of `srcOffset.z` and `extent.depth` does not equal the depth of the subresource specified by `srcSubresource`, `extent.depth` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-vkCmdCopyImage-aspect-06662)VUID-vkCmdCopyImage-aspect-06662  
  If the `aspect` member of any element of `pRegions` includes any flag other than `VK_IMAGE_ASPECT_STENCIL_BIT` or `srcImage` was not created with [separate stencil usage](#VkImageStencilUsageCreateInfo), `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` **must** have been included in the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` used to create `srcImage`
- [](#VUID-vkCmdCopyImage-aspect-06663)VUID-vkCmdCopyImage-aspect-06663  
  If the `aspect` member of any element of `pRegions` includes any flag other than `VK_IMAGE_ASPECT_STENCIL_BIT` or `dstImage` was not created with [separate stencil usage](#VkImageStencilUsageCreateInfo), `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` used to create `dstImage`
- [](#VUID-vkCmdCopyImage-aspect-06664)VUID-vkCmdCopyImage-aspect-06664  
  If the `aspect` member of any element of `pRegions` includes `VK_IMAGE_ASPECT_STENCIL_BIT`, and `srcImage` was created with [separate stencil usage](#VkImageStencilUsageCreateInfo), `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` **must** have been included in the [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage` used to create `srcImage`
- [](#VUID-vkCmdCopyImage-aspect-06665)VUID-vkCmdCopyImage-aspect-06665  
  If the `aspect` member of any element of `pRegions` includes `VK_IMAGE_ASPECT_STENCIL_BIT`, and `dstImage` was created with [separate stencil usage](#VkImageStencilUsageCreateInfo), `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage` used to create `dstImage`

<!--THE END-->

- [](#VUID-vkCmdCopyImage-srcImage-07966)VUID-vkCmdCopyImage-srcImage-07966  
  If `srcImage` is non-sparse then the image or the specified *disjoint* plane **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdCopyImage-srcSubresource-07967)VUID-vkCmdCopyImage-srcSubresource-07967  
  The `srcSubresource.mipLevel` member of each element of `pRegions` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `srcImage` was created
- [](#VUID-vkCmdCopyImage-srcSubresource-07968)VUID-vkCmdCopyImage-srcSubresource-07968  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `srcSubresource.baseArrayLayer` + `srcSubresource.layerCount` of each element of `pRegions` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `srcImage` was created
- [](#VUID-vkCmdCopyImage-srcImage-07969)VUID-vkCmdCopyImage-srcImage-07969  
  `srcImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!--THE END-->

- [](#VUID-vkCmdCopyImage-dstImage-07966)VUID-vkCmdCopyImage-dstImage-07966  
  If `dstImage` is non-sparse then the image or the specified *disjoint* plane **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-vkCmdCopyImage-dstSubresource-07967)VUID-vkCmdCopyImage-dstSubresource-07967  
  The `dstSubresource.mipLevel` member of each element of `pRegions` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-vkCmdCopyImage-dstSubresource-07968)VUID-vkCmdCopyImage-dstSubresource-07968  
  If `dstSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `dstSubresource.baseArrayLayer` + `dstSubresource.layerCount` of each element of `pRegions` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-vkCmdCopyImage-dstImage-07969)VUID-vkCmdCopyImage-dstImage-07969  
  `dstImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyImage-commandBuffer-parameter)VUID-vkCmdCopyImage-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyImage-srcImage-parameter)VUID-vkCmdCopyImage-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkCmdCopyImage-srcImageLayout-parameter)VUID-vkCmdCopyImage-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-vkCmdCopyImage-dstImage-parameter)VUID-vkCmdCopyImage-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-vkCmdCopyImage-dstImageLayout-parameter)VUID-vkCmdCopyImage-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-vkCmdCopyImage-pRegions-parameter)VUID-vkCmdCopyImage-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount` valid [VkImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy.html) structures
- [](#VUID-vkCmdCopyImage-commandBuffer-recording)VUID-vkCmdCopyImage-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyImage-commandBuffer-cmdpool)VUID-vkCmdCopyImage-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, VK\_QUEUE\_GRAPHICS\_BIT, or VK\_QUEUE\_TRANSFER\_BIT operations
- [](#VUID-vkCmdCopyImage-renderpass)VUID-vkCmdCopyImage-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyImage-videocoding)VUID-vkCmdCopyImage-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdCopyImage-regionCount-arraylength)VUID-vkCmdCopyImage-regionCount-arraylength  
  `regionCount` **must** be greater than `0`
- [](#VUID-vkCmdCopyImage-commonparent)VUID-vkCmdCopyImage-commonparent  
  Each of `commandBuffer`, `dstImage`, and `srcImage` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT  
VK\_QUEUE\_TRANSFER\_BIT

Action

Conditional Rendering

vkCmdCopyImage is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageCopy](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyImage).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0