# VkCopyImageInfo2(3) Manual Page

## Name

VkCopyImageInfo2 - Structure specifying parameters of an image copy command



## [](#_c_specification)C Specification

The `VkCopyImageInfo2` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkCopyImageInfo2 {
    VkStructureType        sType;
    const void*            pNext;
    VkImage                srcImage;
    VkImageLayout          srcImageLayout;
    VkImage                dstImage;
    VkImageLayout          dstImageLayout;
    uint32_t               regionCount;
    const VkImageCopy2*    pRegions;
} VkCopyImageInfo2;
```

or the equivalent

```c++
// Provided by VK_KHR_copy_commands2
typedef VkCopyImageInfo2 VkCopyImageInfo2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcImage` is the source image.
- `srcImageLayout` is the current layout of the source image subresource.
- `dstImage` is the destination image.
- `dstImageLayout` is the current layout of the destination image subresource.
- `regionCount` is the number of regions to copy.
- `pRegions` is a pointer to an array of [VkImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy2.html) structures specifying the regions to copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyImageInfo2-pRegions-00124)VUID-VkCopyImageInfo2-pRegions-00124  
  The union of all source regions, and the union of all destination regions, specified by the elements of `pRegions`, **must** not overlap in memory
- [](#VUID-VkCopyImageInfo2-srcImage-01995)VUID-VkCopyImageInfo2-srcImage-01995  
  The [format features](#resources-image-format-features) of `srcImage` **must** contain `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT`
- [](#VUID-VkCopyImageInfo2-srcImageLayout-00128)VUID-VkCopyImageInfo2-srcImageLayout-00128  
  `srcImageLayout` **must** specify the layout of the image subresources of `srcImage` specified in `pRegions` at the time this command is executed on a `VkDevice`
- [](#VUID-VkCopyImageInfo2-srcImageLayout-01917)VUID-VkCopyImageInfo2-srcImageLayout-01917  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-VkCopyImageInfo2-srcImage-09460)VUID-VkCopyImageInfo2-srcImage-09460  
  If `srcImage` and `dstImage` are the same, and any elements of `pRegions` contains the `srcSubresource` and `dstSubresource` with matching `mipLevel` and overlapping array layers, then the `srcImageLayout` and `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`
- [](#VUID-VkCopyImageInfo2-dstImage-01996)VUID-VkCopyImageInfo2-dstImage-01996  
  The [format features](#resources-image-format-features) of `dstImage` **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`
- [](#VUID-VkCopyImageInfo2-dstImageLayout-00133)VUID-VkCopyImageInfo2-dstImageLayout-00133  
  `dstImageLayout` **must** specify the layout of the image subresources of `dstImage` specified in `pRegions` at the time this command is executed on a `VkDevice`
- [](#VUID-VkCopyImageInfo2-dstImageLayout-01395)VUID-VkCopyImageInfo2-dstImageLayout-01395  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`, `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`
- [](#VUID-VkCopyImageInfo2-srcImage-01548)VUID-VkCopyImageInfo2-srcImage-01548  
  If the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each of `srcImage` and `dstImage` is not a [multi-planar format](#formats-multiplanar), the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each of `srcImage` and `dstImage` **must** be [size-compatible](#formats-size-compatibility)
- [](#VUID-VkCopyImageInfo2-None-01549)VUID-VkCopyImageInfo2-None-01549  
  In a copy to or from a plane of a [multi-planar image](#formats-multiplanar), the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of the image and plane **must** be compatible according to [the description of compatible planes](#formats-compatible-planes) for the plane being copied
- [](#VUID-VkCopyImageInfo2-srcImage-09247)VUID-VkCopyImageInfo2-srcImage-09247  
  If the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each of `srcImage` and `dstImage` is a [compressed image format](#compressed_image_formats), the formats **must** have the same texel block extent
- [](#VUID-VkCopyImageInfo2-srcImage-00136)VUID-VkCopyImageInfo2-srcImage-00136  
  The sample count of `srcImage` and `dstImage` **must** match
- [](#VUID-VkCopyImageInfo2-srcOffset-01783)VUID-VkCopyImageInfo2-srcOffset-01783  
  The `srcOffset` and `extent` members of each element of `pRegions` **must** respect the image transfer granularity requirements of `commandBuffer`’s command pool’s queue family, as described in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html)
- [](#VUID-VkCopyImageInfo2-dstOffset-01784)VUID-VkCopyImageInfo2-dstOffset-01784  
  The `dstOffset` and `extent` members of each element of `pRegions` **must** respect the image transfer granularity requirements of `commandBuffer`’s command pool’s queue family, as described in [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties.html)
- [](#VUID-VkCopyImageInfo2-srcImage-01551)VUID-VkCopyImageInfo2-srcImage-01551  
  If neither `srcImage` nor `dstImage` has a [multi-planar format](#formats-multiplanar) and the [`maintenance8`](#features-maintenance8) feature is not enabled then for each element of `pRegions`, `srcSubresource.aspectMask` and `dstSubresource.aspectMask` **must** match
- [](#VUID-VkCopyImageInfo2-srcSubresource-10214)VUID-VkCopyImageInfo2-srcSubresource-10214  
  If `srcSubresource.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT`, then `dstSubresource.aspectMask` **must** not contain both `VK_IMAGE_ASPECT_DEPTH_BIT` and `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkCopyImageInfo2-dstSubresource-10215)VUID-VkCopyImageInfo2-dstSubresource-10215  
  If `dstSubresource.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT`, then `srSubresource.aspectMask` **must** not contain both `VK_IMAGE_ASPECT_DEPTH_BIT` and `VK_IMAGE_ASPECT_STENCIL_BIT`
- [](#VUID-VkCopyImageInfo2-srcImage-08713)VUID-VkCopyImageInfo2-srcImage-08713  
  If `srcImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `srcSubresource.aspectMask` **must** be a single valid [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit
- [](#VUID-VkCopyImageInfo2-dstImage-08714)VUID-VkCopyImageInfo2-dstImage-08714  
  If `dstImage` has a [multi-planar format](#formats-multiplanar), then for each element of `pRegions`, `dstSubresource.aspectMask` **must** be a single valid [multi-planar aspect mask](#formats-multiplanar-image-aspect) bit
- [](#VUID-VkCopyImageInfo2-srcImage-01556)VUID-VkCopyImageInfo2-srcImage-01556  
  If `srcImage` has a [multi-planar format](#formats-multiplanar) and the `dstImage` does not have a multi-planar image format, then for each element of `pRegions`, `dstSubresource.aspectMask` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkCopyImageInfo2-dstImage-01557)VUID-VkCopyImageInfo2-dstImage-01557  
  If `dstImage` has a [multi-planar format](#formats-multiplanar) and the `srcImage` does not have a multi-planar image format, then for each element of `pRegions`, `srcSubresource.aspectMask` **must** be `VK_IMAGE_ASPECT_COLOR_BIT`
- [](#VUID-VkCopyImageInfo2-srcSubresource-10211)VUID-VkCopyImageInfo2-srcSubresource-10211  
  If `srcSubresource.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT` and `dstSubresource.aspectMask` is `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`, then the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values of `srcImage` and `dstImage` **must** be compatible according to [the list of compatible depth-stencil and color formats](#formats-compatible-zs-color)
- [](#VUID-VkCopyImageInfo2-srcSubresource-10212)VUID-VkCopyImageInfo2-srcSubresource-10212  
  If `srcSubresource.aspectMask` is `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT` and `dstSubresource.aspectMask` is `VK_IMAGE_ASPECT_COLOR_BIT`, then the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values of `srcImage` and `dstImage` **must** be compatible according to [the list of compatible depth-stencil and color formats](#formats-compatible-zs-color)
- [](#VUID-VkCopyImageInfo2-apiVersion-07932)VUID-VkCopyImageInfo2-apiVersion-07932  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled, or [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, and either `srcImage` or `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `srcSubresource.baseArrayLayer` and `dstSubresource.baseArrayLayer` **must** both be `0`, and `srcSubresource.layerCount` and `dstSubresource.layerCount` **must** both be `1`
- [](#VUID-VkCopyImageInfo2-srcImage-04443)VUID-VkCopyImageInfo2-srcImage-04443  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `srcSubresource.baseArrayLayer` **must** be `0` and `srcSubresource.layerCount` **must** be `1`
- [](#VUID-VkCopyImageInfo2-dstImage-04444)VUID-VkCopyImageInfo2-dstImage-04444  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `dstSubresource.baseArrayLayer` **must** be `0` and `dstSubresource.layerCount` **must** be `1`
- [](#VUID-VkCopyImageInfo2-aspectMask-00142)VUID-VkCopyImageInfo2-aspectMask-00142  
  For each element of `pRegions`, `srcSubresource.aspectMask` **must** specify aspects present in `srcImage`
- [](#VUID-VkCopyImageInfo2-aspectMask-00143)VUID-VkCopyImageInfo2-aspectMask-00143  
  For each element of `pRegions`, `dstSubresource.aspectMask` **must** specify aspects present in `dstImage`
- [](#VUID-VkCopyImageInfo2-srcOffset-00144)VUID-VkCopyImageInfo2-srcOffset-00144  
  For each element of `pRegions`, `srcOffset.x` and (`extent.width` + `srcOffset.x`) **must** both be greater than or equal to `0` and less than or equal to the width of the specified `srcSubresource` of `srcImage`
- [](#VUID-VkCopyImageInfo2-srcOffset-00145)VUID-VkCopyImageInfo2-srcOffset-00145  
  For each element of `pRegions`, `srcOffset.y` and (`extent.height` + `srcOffset.y`) **must** both be greater than or equal to `0` and less than or equal to the height of the specified `srcSubresource` of `srcImage`
- [](#VUID-VkCopyImageInfo2-srcImage-00146)VUID-VkCopyImageInfo2-srcImage-00146  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `srcOffset.y` **must** be `0` and `extent.height` **must** be `1`
- [](#VUID-VkCopyImageInfo2-srcOffset-00147)VUID-VkCopyImageInfo2-srcOffset-00147  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `srcOffset.z` and (`extent.depth` + `srcOffset.z`) **must** both be greater than or equal to `0` and less than or equal to the depth of the specified `srcSubresource` of `srcImage`
- [](#VUID-VkCopyImageInfo2-srcImage-01785)VUID-VkCopyImageInfo2-srcImage-01785  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `srcOffset.z` **must** be `0` and `extent.depth` **must** be `1`
- [](#VUID-VkCopyImageInfo2-dstImage-01786)VUID-VkCopyImageInfo2-dstImage-01786  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `dstOffset.z` **must** be `0`
- [](#VUID-VkCopyImageInfo2-srcImage-10907)VUID-VkCopyImageInfo2-srcImage-10907  
  If either the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each of `srcImage` and `dstImage` is not a [compressed image format](#compressed_image_formats), and `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `extent.depth` **must** be `1`
- [](#VUID-VkCopyImageInfo2-srcImage-01787)VUID-VkCopyImageInfo2-srcImage-01787  
  If `srcImage` is of type `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `srcOffset.z` **must** be `0`
- [](#VUID-VkCopyImageInfo2-dstImage-01788)VUID-VkCopyImageInfo2-dstImage-01788  
  If `dstImage` is of type `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `dstOffset.z` **must** be `0`
- [](#VUID-VkCopyImageInfo2-apiVersion-07933)VUID-VkCopyImageInfo2-apiVersion-07933  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled, and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, `srcImage` and `dstImage` **must** have the same [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html)
- [](#VUID-VkCopyImageInfo2-apiVersion-08969)VUID-VkCopyImageInfo2-apiVersion-08969  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled, and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, `srcImage` or `dstImage` is of type `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `extent.depth` **must** be `1`
- [](#VUID-VkCopyImageInfo2-srcImage-07743)VUID-VkCopyImageInfo2-srcImage-07743  
  If `srcImage` and `dstImage` have a different [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html), and the [`maintenance5`](#features-maintenance5) feature is not enabled, one **must** be `VK_IMAGE_TYPE_3D` and the other **must** be `VK_IMAGE_TYPE_2D`
- [](#VUID-VkCopyImageInfo2-srcImage-08793)VUID-VkCopyImageInfo2-srcImage-08793  
  If `srcImage` and `dstImage` have the same [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html), for each element of `pRegions`, if neither of the `layerCount` members of `srcSubresource` or `dstSubresource` are `VK_REMAINING_ARRAY_LAYERS`, the `layerCount` members of `srcSubresource` or `dstSubresource` **must** match
- [](#VUID-VkCopyImageInfo2-srcImage-08794)VUID-VkCopyImageInfo2-srcImage-08794  
  If `srcImage` and `dstImage` have the same [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html), and one of the `layerCount` members of `srcSubresource` or `dstSubresource` is `VK_REMAINING_ARRAY_LAYERS`, the other member **must** be either `VK_REMAINING_ARRAY_LAYERS` or equal to the `arrayLayers` member of the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) used to create the image minus `baseArrayLayer`
- [](#VUID-VkCopyImageInfo2-srcImage-01790)VUID-VkCopyImageInfo2-srcImage-01790  
  If `srcImage` and `dstImage` are both of type `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`, `extent.depth` **must** be `1`
- [](#VUID-VkCopyImageInfo2-srcImage-01791)VUID-VkCopyImageInfo2-srcImage-01791  
  If `srcImage` is of type `VK_IMAGE_TYPE_2D`, and `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `extent.depth` **must** equal `srcSubresource.layerCount`
- [](#VUID-VkCopyImageInfo2-dstImage-01792)VUID-VkCopyImageInfo2-dstImage-01792  
  If `dstImage` is of type `VK_IMAGE_TYPE_2D`, and `srcImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `extent.depth` **must** equal `dstSubresource.layerCount`
- [](#VUID-VkCopyImageInfo2-dstOffset-00150)VUID-VkCopyImageInfo2-dstOffset-00150  
  For each element of `pRegions`, `dstOffset.x` and (`extent.width` + `dstOffset.x`), where `extent` is [adjusted for size-compatibility](#formats-size-compatibility), **must** both be greater than or equal to `0` and less than or equal to the width of the specified `dstSubresource` of `dstImage`
- [](#VUID-VkCopyImageInfo2-dstOffset-00151)VUID-VkCopyImageInfo2-dstOffset-00151  
  For each element of `pRegions`, `dstOffset.y` and (`extent.height` + `dstOffset.y`), where `extent` is [adjusted for size-compatibility](#formats-size-compatibility), **must** both be greater than or equal to `0` and less than or equal to the height of the specified `dstSubresource` of `dstImage`
- [](#VUID-VkCopyImageInfo2-dstImage-00152)VUID-VkCopyImageInfo2-dstImage-00152  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `dstOffset.y` **must** be `0`
- [](#VUID-VkCopyImageInfo2-srcImage-10908)VUID-VkCopyImageInfo2-srcImage-10908  
  If either the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of each of `srcImage` and `dstImage` is not a [compressed image format](#compressed_image_formats), and `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of `pRegions`, `extent.height` **must** be `1`, where `extent` is [adjusted for size-compatibility](#formats-size-compatibility)
- [](#VUID-VkCopyImageInfo2-dstOffset-00153)VUID-VkCopyImageInfo2-dstOffset-00153  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`, `dstOffset.z` and (`extent.depth` + `dstOffset.z`), where `extent` is [adjusted for size-compatibility](#formats-size-compatibility), **must** both be greater than or equal to `0` and less than or equal to the depth of the specified `dstSubresource` of `dstImage`
- [](#VUID-VkCopyImageInfo2-pRegions-07278)VUID-VkCopyImageInfo2-pRegions-07278  
  For each element of `pRegions`, `srcOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageInfo2-pRegions-07279)VUID-VkCopyImageInfo2-pRegions-07279  
  For each element of `pRegions`, `srcOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageInfo2-pRegions-07280)VUID-VkCopyImageInfo2-pRegions-07280  
  For each element of `pRegions`, `srcOffset.z` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageInfo2-pRegions-07281)VUID-VkCopyImageInfo2-pRegions-07281  
  For each element of `pRegions`, `dstOffset.x` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyImageInfo2-pRegions-07282)VUID-VkCopyImageInfo2-pRegions-07282  
  For each element of `pRegions`, `dstOffset.y` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyImageInfo2-pRegions-07283)VUID-VkCopyImageInfo2-pRegions-07283  
  For each element of `pRegions`, `dstOffset.z` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `dstImage`
- [](#VUID-VkCopyImageInfo2-srcImage-01728)VUID-VkCopyImageInfo2-srcImage-01728  
  For each element of `pRegions`, if the sum of `srcOffset.x` and `extent.width` does not equal the width of the subresource specified by `srcSubresource`, `extent.width` **must** be a multiple of the [texel block extent width](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageInfo2-srcImage-01729)VUID-VkCopyImageInfo2-srcImage-01729  
  For each element of `pRegions`, if the sum of `srcOffset.y` and `extent.height` does not equal the height of the subresource specified by `srcSubresource`, `extent.height` **must** be a multiple of the [texel block extent height](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageInfo2-srcImage-01730)VUID-VkCopyImageInfo2-srcImage-01730  
  For each element of `pRegions`, if the sum of `srcOffset.z` and `extent.depth` does not equal the depth of the subresource specified by `srcSubresource`, `extent.depth` **must** be a multiple of the [texel block extent depth](#formats-compatibility-classes) of the [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) of `srcImage`
- [](#VUID-VkCopyImageInfo2-aspect-06662)VUID-VkCopyImageInfo2-aspect-06662  
  If the `aspect` member of any element of `pRegions` includes any flag other than `VK_IMAGE_ASPECT_STENCIL_BIT` or `srcImage` was not created with [separate stencil usage](#VkImageStencilUsageCreateInfo), `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` **must** have been included in the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` used to create `srcImage`
- [](#VUID-VkCopyImageInfo2-aspect-06663)VUID-VkCopyImageInfo2-aspect-06663  
  If the `aspect` member of any element of `pRegions` includes any flag other than `VK_IMAGE_ASPECT_STENCIL_BIT` or `dstImage` was not created with [separate stencil usage](#VkImageStencilUsageCreateInfo), `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` used to create `dstImage`
- [](#VUID-VkCopyImageInfo2-aspect-06664)VUID-VkCopyImageInfo2-aspect-06664  
  If the `aspect` member of any element of `pRegions` includes `VK_IMAGE_ASPECT_STENCIL_BIT`, and `srcImage` was created with [separate stencil usage](#VkImageStencilUsageCreateInfo), `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` **must** have been included in the [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage` used to create `srcImage`
- [](#VUID-VkCopyImageInfo2-aspect-06665)VUID-VkCopyImageInfo2-aspect-06665  
  If the `aspect` member of any element of `pRegions` includes `VK_IMAGE_ASPECT_STENCIL_BIT`, and `dstImage` was created with [separate stencil usage](#VkImageStencilUsageCreateInfo), `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage` used to create `dstImage`

<!--THE END-->

- [](#VUID-VkCopyImageInfo2-srcImage-07966)VUID-VkCopyImageInfo2-srcImage-07966  
  If `srcImage` is non-sparse then the image or the specified *disjoint* plane **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkCopyImageInfo2-srcSubresource-07967)VUID-VkCopyImageInfo2-srcSubresource-07967  
  The `srcSubresource.mipLevel` member of each element of `pRegions` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `srcImage` was created
- [](#VUID-VkCopyImageInfo2-srcSubresource-07968)VUID-VkCopyImageInfo2-srcSubresource-07968  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `srcSubresource.baseArrayLayer` + `srcSubresource.layerCount` of each element of `pRegions` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `srcImage` was created
- [](#VUID-VkCopyImageInfo2-srcImage-07969)VUID-VkCopyImageInfo2-srcImage-07969  
  `srcImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!--THE END-->

- [](#VUID-VkCopyImageInfo2-dstImage-07966)VUID-VkCopyImageInfo2-dstImage-07966  
  If `dstImage` is non-sparse then the image or the specified *disjoint* plane **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkCopyImageInfo2-dstSubresource-07967)VUID-VkCopyImageInfo2-dstSubresource-07967  
  The `dstSubresource.mipLevel` member of each element of `pRegions` **must** be less than the `mipLevels` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-VkCopyImageInfo2-dstSubresource-07968)VUID-VkCopyImageInfo2-dstSubresource-07968  
  If `dstSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`, `dstSubresource.baseArrayLayer` + `dstSubresource.layerCount` of each element of `pRegions` **must** be less than or equal to the `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) when `dstImage` was created
- [](#VUID-VkCopyImageInfo2-dstImage-07969)VUID-VkCopyImageInfo2-dstImage-07969  
  `dstImage` **must** not have been created with `flags` containing `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

Valid Usage (Implicit)

- [](#VUID-VkCopyImageInfo2-sType-sType)VUID-VkCopyImageInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_IMAGE_INFO_2`
- [](#VUID-VkCopyImageInfo2-pNext-pNext)VUID-VkCopyImageInfo2-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyImageInfo2-srcImage-parameter)VUID-VkCopyImageInfo2-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkCopyImageInfo2-srcImageLayout-parameter)VUID-VkCopyImageInfo2-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkCopyImageInfo2-dstImage-parameter)VUID-VkCopyImageInfo2-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkCopyImageInfo2-dstImageLayout-parameter)VUID-VkCopyImageInfo2-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html) value
- [](#VUID-VkCopyImageInfo2-pRegions-parameter)VUID-VkCopyImageInfo2-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount` valid [VkImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy2.html) structures
- [](#VUID-VkCopyImageInfo2-regionCount-arraylength)VUID-VkCopyImageInfo2-regionCount-arraylength  
  `regionCount` **must** be greater than `0`
- [](#VUID-VkCopyImageInfo2-commonparent)VUID-VkCopyImageInfo2-commonparent  
  Both of `dstImage`, and `srcImage` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageCopy2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCopy2.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyImage2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImage2.html), [vkCmdCopyImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyImage2KHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyImageInfo2)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0