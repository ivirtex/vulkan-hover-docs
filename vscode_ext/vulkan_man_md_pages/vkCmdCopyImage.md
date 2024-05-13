# vkCmdCopyImage(3) Manual Page

## Name

vkCmdCopyImage - Copy data between images



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data between image objects, call:

``` c
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

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `srcImage` is the source image.

- `srcImageLayout` is the current layout of the source image
  subresource.

- `dstImage` is the destination image.

- `dstImageLayout` is the current layout of the destination image
  subresource.

- `regionCount` is the number of regions to copy.

- `pRegions` is a pointer to an array of [VkImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy.html)
  structures specifying the regions to copy.

## <a href="#_description" class="anchor"></a>Description

Each source region specified by `pRegions` is copied from the source
image to the destination region of the destination image. If any of the
specified regions in `srcImage` overlaps in memory with any of the
specified regions in `dstImage`, values read from those overlapping
regions are undefined.

<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">Multi-planar images</a> **can** only be
copied on a per-plane basis, and the subresources used in each region
when copying to or from such images **must** specify only one plane,
though different regions **can** specify different planes. When copying
planes of multi-planar images, the format considered is the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatible-planes"
target="_blank" rel="noopener">compatible format for that plane</a>,
rather than the format of the multi-planar image.

If the format of the destination image has a different <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-compatibility-classes"
target="_blank" rel="noopener">block extent</a> than the source image
(e.g. one is a compressed format), the offset and extent for each of the
regions specified is <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-size-compatibility"
target="_blank" rel="noopener">scaled according to the block extents of
each format</a> to match in size. Copy regions for each image **must**
be aligned to a multiple of the texel block extent in each dimension,
except at the edges of the image, where region extents **must** match
the edge of the image.

Image data **can** be copied between images with different image types.
If one image is `VK_IMAGE_TYPE_3D` and the other image is
`VK_IMAGE_TYPE_2D` with multiple layers, then each slice is copied to or
from a different layer; `depth` slices in the 3D image correspond to
`layerCount` layers in the 2D image, with an effective `depth` of `1`
used for the 2D image. If <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance5"
target="_blank" rel="noopener"><code>maintenance5</code></a> is enabled,
all other combinations are allowed and function as if 1D images are 2D
images with a height of 1. Otherwise, other combinations of image types
are disallowed.

Valid Usage

- <a href="#VUID-vkCmdCopyImage-commandBuffer-01825"
  id="VUID-vkCmdCopyImage-commandBuffer-01825"></a>
  VUID-vkCmdCopyImage-commandBuffer-01825  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `srcImage` **must** not be a protected image

- <a href="#VUID-vkCmdCopyImage-commandBuffer-01826"
  id="VUID-vkCmdCopyImage-commandBuffer-01826"></a>
  VUID-vkCmdCopyImage-commandBuffer-01826  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be a protected image

- <a href="#VUID-vkCmdCopyImage-commandBuffer-01827"
  id="VUID-vkCmdCopyImage-commandBuffer-01827"></a>
  VUID-vkCmdCopyImage-commandBuffer-01827  
  If `commandBuffer` is a protected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be an unprotected image

<!-- -->

- <a href="#VUID-vkCmdCopyImage-pRegions-00124"
  id="VUID-vkCmdCopyImage-pRegions-00124"></a>
  VUID-vkCmdCopyImage-pRegions-00124  
  The union of all source regions, and the union of all destination
  regions, specified by the elements of `pRegions`, **must** not overlap
  in memory

- <a href="#VUID-vkCmdCopyImage-srcImage-01995"
  id="VUID-vkCmdCopyImage-srcImage-01995"></a>
  VUID-vkCmdCopyImage-srcImage-01995  
  The [format features](#resources-image-format-features) of `srcImage`
  **must** contain `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT`

- <a href="#VUID-vkCmdCopyImage-srcImageLayout-00128"
  id="VUID-vkCmdCopyImage-srcImageLayout-00128"></a>
  VUID-vkCmdCopyImage-srcImageLayout-00128  
  `srcImageLayout` **must** specify the layout of the image subresources
  of `srcImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-vkCmdCopyImage-srcImageLayout-01917"
  id="VUID-vkCmdCopyImage-srcImageLayout-01917"></a>
  VUID-vkCmdCopyImage-srcImageLayout-01917  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-vkCmdCopyImage-srcImage-09460"
  id="VUID-vkCmdCopyImage-srcImage-09460"></a>
  VUID-vkCmdCopyImage-srcImage-09460  
  If `srcImage` and `dstImage` are the same, and any elements of
  `pRegions` contains the `srcSubresource` and `dstSubresource` with
  matching `mipLevel` and overlapping array layers, then the
  `srcImageLayout` and `dstImageLayout` **must** be
  `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`

- <a href="#VUID-vkCmdCopyImage-dstImage-01996"
  id="VUID-vkCmdCopyImage-dstImage-01996"></a>
  VUID-vkCmdCopyImage-dstImage-01996  
  The [format features](#resources-image-format-features) of `dstImage`
  **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`

- <a href="#VUID-vkCmdCopyImage-dstImageLayout-00133"
  id="VUID-vkCmdCopyImage-dstImageLayout-00133"></a>
  VUID-vkCmdCopyImage-dstImageLayout-00133  
  `dstImageLayout` **must** specify the layout of the image subresources
  of `dstImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-vkCmdCopyImage-dstImageLayout-01395"
  id="VUID-vkCmdCopyImage-dstImageLayout-01395"></a>
  VUID-vkCmdCopyImage-dstImageLayout-01395  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-vkCmdCopyImage-srcImage-01548"
  id="VUID-vkCmdCopyImage-srcImage-01548"></a>
  VUID-vkCmdCopyImage-srcImage-01548  
  If the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of each of `srcImage` and `dstImage`
  is not a [*multi-planar
  format*](#formats-requiring-sampler-ycbcr-conversion), the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of each of `srcImage` and `dstImage`
  **must** be [size-compatible](#formats-size-compatibility)

- <a href="#VUID-vkCmdCopyImage-None-01549"
  id="VUID-vkCmdCopyImage-None-01549"></a>
  VUID-vkCmdCopyImage-None-01549  
  In a copy to or from a plane of a [multi-planar
  image](#formats-requiring-sampler-ycbcr-conversion), the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of the image and plane **must** be
  compatible according to [the description of compatible
  planes](#formats-compatible-planes) for the plane being copied

- <a href="#VUID-vkCmdCopyImage-srcImage-09247"
  id="VUID-vkCmdCopyImage-srcImage-09247"></a>
  VUID-vkCmdCopyImage-srcImage-09247  
  If the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of each of `srcImage` and `dstImage`
  is a [compressed image format](#compressed_image_formats), the formats
  **must** have the same texel block extent

- <a href="#VUID-vkCmdCopyImage-srcImage-00136"
  id="VUID-vkCmdCopyImage-srcImage-00136"></a>
  VUID-vkCmdCopyImage-srcImage-00136  
  The sample count of `srcImage` and `dstImage` **must** match

- <a href="#VUID-vkCmdCopyImage-srcOffset-01783"
  id="VUID-vkCmdCopyImage-srcOffset-01783"></a>
  VUID-vkCmdCopyImage-srcOffset-01783  
  The `srcOffset` and `extent` members of each element of `pRegions`
  **must** respect the image transfer granularity requirements of
  `commandBuffer`’s command pool’s queue family, as described in
  [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html)

- <a href="#VUID-vkCmdCopyImage-dstOffset-01784"
  id="VUID-vkCmdCopyImage-dstOffset-01784"></a>
  VUID-vkCmdCopyImage-dstOffset-01784  
  The `dstOffset` and `extent` members of each element of `pRegions`
  **must** respect the image transfer granularity requirements of
  `commandBuffer`’s command pool’s queue family, as described in
  [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html)

- <a href="#VUID-vkCmdCopyImage-srcImage-01551"
  id="VUID-vkCmdCopyImage-srcImage-01551"></a>
  VUID-vkCmdCopyImage-srcImage-01551  
  If neither `srcImage` nor `dstImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion) then for each
  element of `pRegions`, `srcSubresource.aspectMask` and
  `dstSubresource.aspectMask` **must** match

- <a href="#VUID-vkCmdCopyImage-srcImage-08713"
  id="VUID-vkCmdCopyImage-srcImage-08713"></a>
  VUID-vkCmdCopyImage-srcImage-08713  
  If `srcImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `srcSubresource.aspectMask` **must** be a
  single valid [multi-planar aspect mask](#formats-planes-image-aspect)
  bit

- <a href="#VUID-vkCmdCopyImage-dstImage-08714"
  id="VUID-vkCmdCopyImage-dstImage-08714"></a>
  VUID-vkCmdCopyImage-dstImage-08714  
  If `dstImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `dstSubresource.aspectMask` **must** be a
  single valid [multi-planar aspect mask](#formats-planes-image-aspect)
  bit

- <a href="#VUID-vkCmdCopyImage-srcImage-01556"
  id="VUID-vkCmdCopyImage-srcImage-01556"></a>
  VUID-vkCmdCopyImage-srcImage-01556  
  If `srcImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion) and the
  `dstImage` does not have a multi-planar image format, then for each
  element of `pRegions`, `dstSubresource.aspectMask` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-vkCmdCopyImage-dstImage-01557"
  id="VUID-vkCmdCopyImage-dstImage-01557"></a>
  VUID-vkCmdCopyImage-dstImage-01557  
  If `dstImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion) and the
  `srcImage` does not have a multi-planar image format, then for each
  element of `pRegions`, `srcSubresource.aspectMask` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-vkCmdCopyImage-apiVersion-07932"
  id="VUID-vkCmdCopyImage-apiVersion-07932"></a>
  VUID-vkCmdCopyImage-apiVersion-07932  
  If the [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html) extension is
  not enabled, or
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, and either `srcImage` or `dstImage` is of
  type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`,
  `srcSubresource.baseArrayLayer` and `dstSubresource.baseArrayLayer`
  **must** both be `0`, and `srcSubresource.layerCount` and
  `dstSubresource.layerCount` **must** both be `1`

- <a href="#VUID-vkCmdCopyImage-srcImage-04443"
  id="VUID-vkCmdCopyImage-srcImage-04443"></a>
  VUID-vkCmdCopyImage-srcImage-04443  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `srcSubresource.baseArrayLayer` **must** be `0` and
  `srcSubresource.layerCount` **must** be `1`

- <a href="#VUID-vkCmdCopyImage-dstImage-04444"
  id="VUID-vkCmdCopyImage-dstImage-04444"></a>
  VUID-vkCmdCopyImage-dstImage-04444  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `dstSubresource.baseArrayLayer` **must** be `0` and
  `dstSubresource.layerCount` **must** be `1`

- <a href="#VUID-vkCmdCopyImage-aspectMask-00142"
  id="VUID-vkCmdCopyImage-aspectMask-00142"></a>
  VUID-vkCmdCopyImage-aspectMask-00142  
  For each element of `pRegions`, `srcSubresource.aspectMask` **must**
  specify aspects present in `srcImage`

- <a href="#VUID-vkCmdCopyImage-aspectMask-00143"
  id="VUID-vkCmdCopyImage-aspectMask-00143"></a>
  VUID-vkCmdCopyImage-aspectMask-00143  
  For each element of `pRegions`, `dstSubresource.aspectMask` **must**
  specify aspects present in `dstImage`

- <a href="#VUID-vkCmdCopyImage-srcOffset-00144"
  id="VUID-vkCmdCopyImage-srcOffset-00144"></a>
  VUID-vkCmdCopyImage-srcOffset-00144  
  For each element of `pRegions`, `srcOffset.x` and (`extent.width` +
  `srcOffset.x`) **must** both be greater than or equal to `0` and less
  than or equal to the width of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-vkCmdCopyImage-srcOffset-00145"
  id="VUID-vkCmdCopyImage-srcOffset-00145"></a>
  VUID-vkCmdCopyImage-srcOffset-00145  
  For each element of `pRegions`, `srcOffset.y` and (`extent.height` +
  `srcOffset.y`) **must** both be greater than or equal to `0` and less
  than or equal to the height of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-vkCmdCopyImage-srcImage-00146"
  id="VUID-vkCmdCopyImage-srcImage-00146"></a>
  VUID-vkCmdCopyImage-srcImage-00146  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `srcOffset.y` **must** be `0` and `extent.height` **must**
  be `1`

- <a href="#VUID-vkCmdCopyImage-srcOffset-00147"
  id="VUID-vkCmdCopyImage-srcOffset-00147"></a>
  VUID-vkCmdCopyImage-srcOffset-00147  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `srcOffset.z` and (`extent.depth` + `srcOffset.z`)
  **must** both be greater than or equal to `0` and less than or equal
  to the depth of the specified `srcSubresource` of `srcImage`

- <a href="#VUID-vkCmdCopyImage-srcImage-01785"
  id="VUID-vkCmdCopyImage-srcImage-01785"></a>
  VUID-vkCmdCopyImage-srcImage-01785  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `srcOffset.z` **must** be `0` and `extent.depth` **must**
  be `1`

- <a href="#VUID-vkCmdCopyImage-dstImage-01786"
  id="VUID-vkCmdCopyImage-dstImage-01786"></a>
  VUID-vkCmdCopyImage-dstImage-01786  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `dstOffset.z` **must** be `0` and `extent.depth` **must**
  be `1`

- <a href="#VUID-vkCmdCopyImage-srcImage-01787"
  id="VUID-vkCmdCopyImage-srcImage-01787"></a>
  VUID-vkCmdCopyImage-srcImage-01787  
  If `srcImage` is of type `VK_IMAGE_TYPE_2D`, then for each element of
  `pRegions`, `srcOffset.z` **must** be `0`

- <a href="#VUID-vkCmdCopyImage-dstImage-01788"
  id="VUID-vkCmdCopyImage-dstImage-01788"></a>
  VUID-vkCmdCopyImage-dstImage-01788  
  If `dstImage` is of type `VK_IMAGE_TYPE_2D`, then for each element of
  `pRegions`, `dstOffset.z` **must** be `0`

- <a href="#VUID-vkCmdCopyImage-apiVersion-07933"
  id="VUID-vkCmdCopyImage-apiVersion-07933"></a>
  VUID-vkCmdCopyImage-apiVersion-07933  
  If the [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html) extension is
  not enabled, and
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, `srcImage` and `dstImage` **must** have the
  same [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html)

- <a href="#VUID-vkCmdCopyImage-apiVersion-08969"
  id="VUID-vkCmdCopyImage-apiVersion-08969"></a>
  VUID-vkCmdCopyImage-apiVersion-08969  
  If the [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html) extension is
  not enabled, and
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, `srcImage` or `dstImage` is of type
  `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`,
  `extent.depth` **must** be `1`

- <a href="#VUID-vkCmdCopyImage-srcImage-07743"
  id="VUID-vkCmdCopyImage-srcImage-07743"></a>
  VUID-vkCmdCopyImage-srcImage-07743  
  If `srcImage` and `dstImage` have a different
  [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html), and
  [`maintenance5`](#features-maintenance5) is not enabled, one **must**
  be `VK_IMAGE_TYPE_3D` and the other **must** be `VK_IMAGE_TYPE_2D`

- <a href="#VUID-vkCmdCopyImage-srcImage-08793"
  id="VUID-vkCmdCopyImage-srcImage-08793"></a>
  VUID-vkCmdCopyImage-srcImage-08793  
  If `srcImage` and `dstImage` have the same
  [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html), for each element of `pRegions`, if
  neither of the `layerCount` members of `srcSubresource` or
  `dstSubresource` are `VK_REMAINING_ARRAY_LAYERS`, the `layerCount`
  members of `srcSubresource` or `dstSubresource` **must** match

- <a href="#VUID-vkCmdCopyImage-srcImage-08794"
  id="VUID-vkCmdCopyImage-srcImage-08794"></a>
  VUID-vkCmdCopyImage-srcImage-08794  
  If `srcImage` and `dstImage` have the same
  [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html), and one of the `layerCount` members
  of `srcSubresource` or `dstSubresource` is
  `VK_REMAINING_ARRAY_LAYERS`, the other member **must** be either
  `VK_REMAINING_ARRAY_LAYERS` or equal to the `arrayLayers` member of
  the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) used to create the
  image minus `baseArrayLayer`

- <a href="#VUID-vkCmdCopyImage-srcImage-01790"
  id="VUID-vkCmdCopyImage-srcImage-01790"></a>
  VUID-vkCmdCopyImage-srcImage-01790  
  If `srcImage` and `dstImage` are both of type `VK_IMAGE_TYPE_2D`, then
  for each element of `pRegions`, `extent.depth` **must** be `1`

- <a href="#VUID-vkCmdCopyImage-srcImage-01791"
  id="VUID-vkCmdCopyImage-srcImage-01791"></a>
  VUID-vkCmdCopyImage-srcImage-01791  
  If `srcImage` is of type `VK_IMAGE_TYPE_2D`, and `dstImage` is of type
  `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`,
  `extent.depth` **must** equal `srcSubresource.layerCount`

- <a href="#VUID-vkCmdCopyImage-dstImage-01792"
  id="VUID-vkCmdCopyImage-dstImage-01792"></a>
  VUID-vkCmdCopyImage-dstImage-01792  
  If `dstImage` is of type `VK_IMAGE_TYPE_2D`, and `srcImage` is of type
  `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`,
  `extent.depth` **must** equal `dstSubresource.layerCount`

- <a href="#VUID-vkCmdCopyImage-dstOffset-00150"
  id="VUID-vkCmdCopyImage-dstOffset-00150"></a>
  VUID-vkCmdCopyImage-dstOffset-00150  
  For each element of `pRegions`, `dstOffset.x` and (`extent.width` +
  `dstOffset.x`) **must** both be greater than or equal to `0` and less
  than or equal to the width of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-vkCmdCopyImage-dstOffset-00151"
  id="VUID-vkCmdCopyImage-dstOffset-00151"></a>
  VUID-vkCmdCopyImage-dstOffset-00151  
  For each element of `pRegions`, `dstOffset.y` and (`extent.height` +
  `dstOffset.y`) **must** both be greater than or equal to `0` and less
  than or equal to the height of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-vkCmdCopyImage-dstImage-00152"
  id="VUID-vkCmdCopyImage-dstImage-00152"></a>
  VUID-vkCmdCopyImage-dstImage-00152  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `dstOffset.y` **must** be `0` and `extent.height` **must**
  be `1`

- <a href="#VUID-vkCmdCopyImage-dstOffset-00153"
  id="VUID-vkCmdCopyImage-dstOffset-00153"></a>
  VUID-vkCmdCopyImage-dstOffset-00153  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `dstOffset.z` and (`extent.depth` + `dstOffset.z`)
  **must** both be greater than or equal to `0` and less than or equal
  to the depth of the specified `dstSubresource` of `dstImage`

- <a href="#VUID-vkCmdCopyImage-pRegions-07278"
  id="VUID-vkCmdCopyImage-pRegions-07278"></a>
  VUID-vkCmdCopyImage-pRegions-07278  
  For each element of `pRegions`, `srcOffset.x` **must** be a multiple
  of the [texel block extent width](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImage-pRegions-07279"
  id="VUID-vkCmdCopyImage-pRegions-07279"></a>
  VUID-vkCmdCopyImage-pRegions-07279  
  For each element of `pRegions`, `srcOffset.y` **must** be a multiple
  of the [texel block extent height](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImage-pRegions-07280"
  id="VUID-vkCmdCopyImage-pRegions-07280"></a>
  VUID-vkCmdCopyImage-pRegions-07280  
  For each element of `pRegions`, `srcOffset.z` **must** be a multiple
  of the [texel block extent depth](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImage-pRegions-07281"
  id="VUID-vkCmdCopyImage-pRegions-07281"></a>
  VUID-vkCmdCopyImage-pRegions-07281  
  For each element of `pRegions`, `dstOffset.x` **must** be a multiple
  of the [texel block extent width](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyImage-pRegions-07282"
  id="VUID-vkCmdCopyImage-pRegions-07282"></a>
  VUID-vkCmdCopyImage-pRegions-07282  
  For each element of `pRegions`, `dstOffset.y` **must** be a multiple
  of the [texel block extent height](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyImage-pRegions-07283"
  id="VUID-vkCmdCopyImage-pRegions-07283"></a>
  VUID-vkCmdCopyImage-pRegions-07283  
  For each element of `pRegions`, `dstOffset.z` **must** be a multiple
  of the [texel block extent depth](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyImage-srcImage-01728"
  id="VUID-vkCmdCopyImage-srcImage-01728"></a>
  VUID-vkCmdCopyImage-srcImage-01728  
  For each element of `pRegions`, if the sum of `srcOffset.x` and
  `extent.width` does not equal the width of the subresource specified
  by `srcSubresource`, `extent.width` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImage-srcImage-01729"
  id="VUID-vkCmdCopyImage-srcImage-01729"></a>
  VUID-vkCmdCopyImage-srcImage-01729  
  For each element of `pRegions`, if the sum of `srcOffset.y` and
  `extent.height` does not equal the height of the subresource specified
  by `srcSubresource`, `extent.height` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImage-srcImage-01730"
  id="VUID-vkCmdCopyImage-srcImage-01730"></a>
  VUID-vkCmdCopyImage-srcImage-01730  
  For each element of `pRegions`, if the sum of `srcOffset.z` and
  `extent.depth` does not equal the depth of the subresource specified
  by `srcSubresource`, `extent.depth` **must** be a multiple of the
  [texel block extent depth](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-vkCmdCopyImage-dstImage-01732"
  id="VUID-vkCmdCopyImage-dstImage-01732"></a>
  VUID-vkCmdCopyImage-dstImage-01732  
  For each element of `pRegions`, if the sum of `dstOffset.x` and
  `extent.width` does not equal the width of the subresource specified
  by `dstSubresource`, `extent.width` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyImage-dstImage-01733"
  id="VUID-vkCmdCopyImage-dstImage-01733"></a>
  VUID-vkCmdCopyImage-dstImage-01733  
  For each element of `pRegions`, if the sum of `dstOffset.y` and
  `extent.height` does not equal the height of the subresource specified
  by `dstSubresource`, `extent.height` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyImage-dstImage-01734"
  id="VUID-vkCmdCopyImage-dstImage-01734"></a>
  VUID-vkCmdCopyImage-dstImage-01734  
  For each element of `pRegions`, if the sum of `dstOffset.z` and
  `extent.depth` does not equal the depth of the subresource specified
  by `dstSubresource`, `extent.depth` **must** be a multiple of the
  [texel block extent depth](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-vkCmdCopyImage-aspect-06662"
  id="VUID-vkCmdCopyImage-aspect-06662"></a>
  VUID-vkCmdCopyImage-aspect-06662  
  If the `aspect` member of any element of `pRegions` includes any flag
  other than `VK_IMAGE_ASPECT_STENCIL_BIT` or `srcImage` was not created
  with [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` **must** have been included in the
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` used to create
  `srcImage`

- <a href="#VUID-vkCmdCopyImage-aspect-06663"
  id="VUID-vkCmdCopyImage-aspect-06663"></a>
  VUID-vkCmdCopyImage-aspect-06663  
  If the `aspect` member of any element of `pRegions` includes any flag
  other than `VK_IMAGE_ASPECT_STENCIL_BIT` or `dstImage` was not created
  with [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` used to create
  `dstImage`

- <a href="#VUID-vkCmdCopyImage-aspect-06664"
  id="VUID-vkCmdCopyImage-aspect-06664"></a>
  VUID-vkCmdCopyImage-aspect-06664  
  If the `aspect` member of any element of `pRegions` includes
  `VK_IMAGE_ASPECT_STENCIL_BIT`, and `srcImage` was created with
  [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` **must** have been included in the
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`
  used to create `srcImage`

- <a href="#VUID-vkCmdCopyImage-aspect-06665"
  id="VUID-vkCmdCopyImage-aspect-06665"></a>
  VUID-vkCmdCopyImage-aspect-06665  
  If the `aspect` member of any element of `pRegions` includes
  `VK_IMAGE_ASPECT_STENCIL_BIT`, and `dstImage` was created with
  [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`
  used to create `dstImage`

<!-- -->

- <a href="#VUID-vkCmdCopyImage-srcImage-07966"
  id="VUID-vkCmdCopyImage-srcImage-07966"></a>
  VUID-vkCmdCopyImage-srcImage-07966  
  If `srcImage` is non-sparse then the image or the specified *disjoint*
  plane **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-vkCmdCopyImage-srcSubresource-07967"
  id="VUID-vkCmdCopyImage-srcSubresource-07967"></a>
  VUID-vkCmdCopyImage-srcSubresource-07967  
  The `srcSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `srcImage` was
  created

- <a href="#VUID-vkCmdCopyImage-srcSubresource-07968"
  id="VUID-vkCmdCopyImage-srcSubresource-07968"></a>
  VUID-vkCmdCopyImage-srcSubresource-07968  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `srcSubresource.baseArrayLayer` + `srcSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `srcImage` was created

- <a href="#VUID-vkCmdCopyImage-srcImage-07969"
  id="VUID-vkCmdCopyImage-srcImage-07969"></a>
  VUID-vkCmdCopyImage-srcImage-07969  
  `srcImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!-- -->

- <a href="#VUID-vkCmdCopyImage-dstImage-07966"
  id="VUID-vkCmdCopyImage-dstImage-07966"></a>
  VUID-vkCmdCopyImage-dstImage-07966  
  If `dstImage` is non-sparse then the image or the specified *disjoint*
  plane **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-vkCmdCopyImage-dstSubresource-07967"
  id="VUID-vkCmdCopyImage-dstSubresource-07967"></a>
  VUID-vkCmdCopyImage-dstSubresource-07967  
  The `dstSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `dstImage` was
  created

- <a href="#VUID-vkCmdCopyImage-dstSubresource-07968"
  id="VUID-vkCmdCopyImage-dstSubresource-07968"></a>
  VUID-vkCmdCopyImage-dstSubresource-07968  
  If `dstSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `dstSubresource.baseArrayLayer` + `dstSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `dstImage` was created

- <a href="#VUID-vkCmdCopyImage-dstImage-07969"
  id="VUID-vkCmdCopyImage-dstImage-07969"></a>
  VUID-vkCmdCopyImage-dstImage-07969  
  `dstImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdCopyImage-commandBuffer-parameter"
  id="VUID-vkCmdCopyImage-commandBuffer-parameter"></a>
  VUID-vkCmdCopyImage-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdCopyImage-srcImage-parameter"
  id="VUID-vkCmdCopyImage-srcImage-parameter"></a>
  VUID-vkCmdCopyImage-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkCmdCopyImage-srcImageLayout-parameter"
  id="VUID-vkCmdCopyImage-srcImageLayout-parameter"></a>
  VUID-vkCmdCopyImage-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-vkCmdCopyImage-dstImage-parameter"
  id="VUID-vkCmdCopyImage-dstImage-parameter"></a>
  VUID-vkCmdCopyImage-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkCmdCopyImage-dstImageLayout-parameter"
  id="VUID-vkCmdCopyImage-dstImageLayout-parameter"></a>
  VUID-vkCmdCopyImage-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-vkCmdCopyImage-pRegions-parameter"
  id="VUID-vkCmdCopyImage-pRegions-parameter"></a>
  VUID-vkCmdCopyImage-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy.html) structures

- <a href="#VUID-vkCmdCopyImage-commandBuffer-recording"
  id="VUID-vkCmdCopyImage-commandBuffer-recording"></a>
  VUID-vkCmdCopyImage-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdCopyImage-commandBuffer-cmdpool"
  id="VUID-vkCmdCopyImage-commandBuffer-cmdpool"></a>
  VUID-vkCmdCopyImage-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support transfer, graphics, or compute operations

- <a href="#VUID-vkCmdCopyImage-renderpass"
  id="VUID-vkCmdCopyImage-renderpass"></a>
  VUID-vkCmdCopyImage-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdCopyImage-videocoding"
  id="VUID-vkCmdCopyImage-videocoding"></a>
  VUID-vkCmdCopyImage-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdCopyImage-regionCount-arraylength"
  id="VUID-vkCmdCopyImage-regionCount-arraylength"></a>
  VUID-vkCmdCopyImage-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-vkCmdCopyImage-commonparent"
  id="VUID-vkCmdCopyImage-commonparent"></a>
  VUID-vkCmdCopyImage-commonparent  
  Each of `commandBuffer`, `dstImage`, and `srcImage` **must** have been
  created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr class="header">
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Transfer<br />
Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageCopy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdCopyImage"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
