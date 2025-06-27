# VkCopyImageInfo2(3) Manual Page

## Name

VkCopyImageInfo2 - Structure specifying parameters of an image copy
command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyImageInfo2` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_copy_commands2
typedef VkCopyImageInfo2 VkCopyImageInfo2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcImage` is the source image.

- `srcImageLayout` is the current layout of the source image
  subresource.

- `dstImage` is the destination image.

- `dstImageLayout` is the current layout of the destination image
  subresource.

- `regionCount` is the number of regions to copy.

- `pRegions` is a pointer to an array of
  [VkImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy2.html) structures specifying the regions to
  copy.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCopyImageInfo2-pRegions-00124"
  id="VUID-VkCopyImageInfo2-pRegions-00124"></a>
  VUID-VkCopyImageInfo2-pRegions-00124  
  The union of all source regions, and the union of all destination
  regions, specified by the elements of `pRegions`, **must** not overlap
  in memory

- <a href="#VUID-VkCopyImageInfo2-srcImage-01995"
  id="VUID-VkCopyImageInfo2-srcImage-01995"></a>
  VUID-VkCopyImageInfo2-srcImage-01995  
  The [format features](#resources-image-format-features) of `srcImage`
  **must** contain `VK_FORMAT_FEATURE_TRANSFER_SRC_BIT`

- <a href="#VUID-VkCopyImageInfo2-srcImageLayout-00128"
  id="VUID-VkCopyImageInfo2-srcImageLayout-00128"></a>
  VUID-VkCopyImageInfo2-srcImageLayout-00128  
  `srcImageLayout` **must** specify the layout of the image subresources
  of `srcImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-VkCopyImageInfo2-srcImageLayout-01917"
  id="VUID-VkCopyImageInfo2-srcImageLayout-01917"></a>
  VUID-VkCopyImageInfo2-srcImageLayout-01917  
  `srcImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-VkCopyImageInfo2-srcImage-09460"
  id="VUID-VkCopyImageInfo2-srcImage-09460"></a>
  VUID-VkCopyImageInfo2-srcImage-09460  
  If `srcImage` and `dstImage` are the same, and any elements of
  `pRegions` contains the `srcSubresource` and `dstSubresource` with
  matching `mipLevel` and overlapping array layers, then the
  `srcImageLayout` and `dstImageLayout` **must** be
  `VK_IMAGE_LAYOUT_GENERAL` or `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`

- <a href="#VUID-VkCopyImageInfo2-dstImage-01996"
  id="VUID-VkCopyImageInfo2-dstImage-01996"></a>
  VUID-VkCopyImageInfo2-dstImage-01996  
  The [format features](#resources-image-format-features) of `dstImage`
  **must** contain `VK_FORMAT_FEATURE_TRANSFER_DST_BIT`

- <a href="#VUID-VkCopyImageInfo2-dstImageLayout-00133"
  id="VUID-VkCopyImageInfo2-dstImageLayout-00133"></a>
  VUID-VkCopyImageInfo2-dstImageLayout-00133  
  `dstImageLayout` **must** specify the layout of the image subresources
  of `dstImage` specified in `pRegions` at the time this command is
  executed on a `VkDevice`

- <a href="#VUID-VkCopyImageInfo2-dstImageLayout-01395"
  id="VUID-VkCopyImageInfo2-dstImageLayout-01395"></a>
  VUID-VkCopyImageInfo2-dstImageLayout-01395  
  `dstImageLayout` **must** be `VK_IMAGE_LAYOUT_SHARED_PRESENT_KHR`,
  `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL`, or `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-VkCopyImageInfo2-srcImage-01548"
  id="VUID-VkCopyImageInfo2-srcImage-01548"></a>
  VUID-VkCopyImageInfo2-srcImage-01548  
  If the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of each of `srcImage` and `dstImage`
  is not a [*multi-planar
  format*](#formats-requiring-sampler-ycbcr-conversion), the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of each of `srcImage` and `dstImage`
  **must** be [size-compatible](#formats-size-compatibility)

- <a href="#VUID-VkCopyImageInfo2-None-01549"
  id="VUID-VkCopyImageInfo2-None-01549"></a>
  VUID-VkCopyImageInfo2-None-01549  
  In a copy to or from a plane of a [multi-planar
  image](#formats-requiring-sampler-ycbcr-conversion), the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of the image and plane **must** be
  compatible according to [the description of compatible
  planes](#formats-compatible-planes) for the plane being copied

- <a href="#VUID-VkCopyImageInfo2-srcImage-09247"
  id="VUID-VkCopyImageInfo2-srcImage-09247"></a>
  VUID-VkCopyImageInfo2-srcImage-09247  
  If the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of each of `srcImage` and `dstImage`
  is a [compressed image format](#compressed_image_formats), the formats
  **must** have the same texel block extent

- <a href="#VUID-VkCopyImageInfo2-srcImage-00136"
  id="VUID-VkCopyImageInfo2-srcImage-00136"></a>
  VUID-VkCopyImageInfo2-srcImage-00136  
  The sample count of `srcImage` and `dstImage` **must** match

- <a href="#VUID-VkCopyImageInfo2-srcOffset-01783"
  id="VUID-VkCopyImageInfo2-srcOffset-01783"></a>
  VUID-VkCopyImageInfo2-srcOffset-01783  
  The `srcOffset` and `extent` members of each element of `pRegions`
  **must** respect the image transfer granularity requirements of
  `commandBuffer`’s command pool’s queue family, as described in
  [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html)

- <a href="#VUID-VkCopyImageInfo2-dstOffset-01784"
  id="VUID-VkCopyImageInfo2-dstOffset-01784"></a>
  VUID-VkCopyImageInfo2-dstOffset-01784  
  The `dstOffset` and `extent` members of each element of `pRegions`
  **must** respect the image transfer granularity requirements of
  `commandBuffer`’s command pool’s queue family, as described in
  [VkQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties.html)

- <a href="#VUID-VkCopyImageInfo2-srcImage-01551"
  id="VUID-VkCopyImageInfo2-srcImage-01551"></a>
  VUID-VkCopyImageInfo2-srcImage-01551  
  If neither `srcImage` nor `dstImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion) then for each
  element of `pRegions`, `srcSubresource.aspectMask` and
  `dstSubresource.aspectMask` **must** match

- <a href="#VUID-VkCopyImageInfo2-srcImage-08713"
  id="VUID-VkCopyImageInfo2-srcImage-08713"></a>
  VUID-VkCopyImageInfo2-srcImage-08713  
  If `srcImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `srcSubresource.aspectMask` **must** be a
  single valid [multi-planar aspect mask](#formats-planes-image-aspect)
  bit

- <a href="#VUID-VkCopyImageInfo2-dstImage-08714"
  id="VUID-VkCopyImageInfo2-dstImage-08714"></a>
  VUID-VkCopyImageInfo2-dstImage-08714  
  If `dstImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then for each
  element of `pRegions`, `dstSubresource.aspectMask` **must** be a
  single valid [multi-planar aspect mask](#formats-planes-image-aspect)
  bit

- <a href="#VUID-VkCopyImageInfo2-srcImage-01556"
  id="VUID-VkCopyImageInfo2-srcImage-01556"></a>
  VUID-VkCopyImageInfo2-srcImage-01556  
  If `srcImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion) and the
  `dstImage` does not have a multi-planar image format, then for each
  element of `pRegions`, `dstSubresource.aspectMask` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkCopyImageInfo2-dstImage-01557"
  id="VUID-VkCopyImageInfo2-dstImage-01557"></a>
  VUID-VkCopyImageInfo2-dstImage-01557  
  If `dstImage` has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion) and the
  `srcImage` does not have a multi-planar image format, then for each
  element of `pRegions`, `srcSubresource.aspectMask` **must** be
  `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkCopyImageInfo2-apiVersion-07932"
  id="VUID-VkCopyImageInfo2-apiVersion-07932"></a>
  VUID-VkCopyImageInfo2-apiVersion-07932  
  If the [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html) extension is
  not enabled, or
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, and either `srcImage` or `dstImage` is of
  type `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`,
  `srcSubresource.baseArrayLayer` and `dstSubresource.baseArrayLayer`
  **must** both be `0`, and `srcSubresource.layerCount` and
  `dstSubresource.layerCount` **must** both be `1`

- <a href="#VUID-VkCopyImageInfo2-srcImage-04443"
  id="VUID-VkCopyImageInfo2-srcImage-04443"></a>
  VUID-VkCopyImageInfo2-srcImage-04443  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `srcSubresource.baseArrayLayer` **must** be `0` and
  `srcSubresource.layerCount` **must** be `1`

- <a href="#VUID-VkCopyImageInfo2-dstImage-04444"
  id="VUID-VkCopyImageInfo2-dstImage-04444"></a>
  VUID-VkCopyImageInfo2-dstImage-04444  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `dstSubresource.baseArrayLayer` **must** be `0` and
  `dstSubresource.layerCount` **must** be `1`

- <a href="#VUID-VkCopyImageInfo2-aspectMask-00142"
  id="VUID-VkCopyImageInfo2-aspectMask-00142"></a>
  VUID-VkCopyImageInfo2-aspectMask-00142  
  For each element of `pRegions`, `srcSubresource.aspectMask` **must**
  specify aspects present in `srcImage`

- <a href="#VUID-VkCopyImageInfo2-aspectMask-00143"
  id="VUID-VkCopyImageInfo2-aspectMask-00143"></a>
  VUID-VkCopyImageInfo2-aspectMask-00143  
  For each element of `pRegions`, `dstSubresource.aspectMask` **must**
  specify aspects present in `dstImage`

- <a href="#VUID-VkCopyImageInfo2-srcOffset-00144"
  id="VUID-VkCopyImageInfo2-srcOffset-00144"></a>
  VUID-VkCopyImageInfo2-srcOffset-00144  
  For each element of `pRegions`, `srcOffset.x` and (`extent.width` +
  `srcOffset.x`) **must** both be greater than or equal to `0` and less
  than or equal to the width of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-VkCopyImageInfo2-srcOffset-00145"
  id="VUID-VkCopyImageInfo2-srcOffset-00145"></a>
  VUID-VkCopyImageInfo2-srcOffset-00145  
  For each element of `pRegions`, `srcOffset.y` and (`extent.height` +
  `srcOffset.y`) **must** both be greater than or equal to `0` and less
  than or equal to the height of the specified `srcSubresource` of
  `srcImage`

- <a href="#VUID-VkCopyImageInfo2-srcImage-00146"
  id="VUID-VkCopyImageInfo2-srcImage-00146"></a>
  VUID-VkCopyImageInfo2-srcImage-00146  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `srcOffset.y` **must** be `0` and `extent.height` **must**
  be `1`

- <a href="#VUID-VkCopyImageInfo2-srcOffset-00147"
  id="VUID-VkCopyImageInfo2-srcOffset-00147"></a>
  VUID-VkCopyImageInfo2-srcOffset-00147  
  If `srcImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `srcOffset.z` and (`extent.depth` + `srcOffset.z`)
  **must** both be greater than or equal to `0` and less than or equal
  to the depth of the specified `srcSubresource` of `srcImage`

- <a href="#VUID-VkCopyImageInfo2-srcImage-01785"
  id="VUID-VkCopyImageInfo2-srcImage-01785"></a>
  VUID-VkCopyImageInfo2-srcImage-01785  
  If `srcImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `srcOffset.z` **must** be `0` and `extent.depth` **must**
  be `1`

- <a href="#VUID-VkCopyImageInfo2-dstImage-01786"
  id="VUID-VkCopyImageInfo2-dstImage-01786"></a>
  VUID-VkCopyImageInfo2-dstImage-01786  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `dstOffset.z` **must** be `0` and `extent.depth` **must**
  be `1`

- <a href="#VUID-VkCopyImageInfo2-srcImage-01787"
  id="VUID-VkCopyImageInfo2-srcImage-01787"></a>
  VUID-VkCopyImageInfo2-srcImage-01787  
  If `srcImage` is of type `VK_IMAGE_TYPE_2D`, then for each element of
  `pRegions`, `srcOffset.z` **must** be `0`

- <a href="#VUID-VkCopyImageInfo2-dstImage-01788"
  id="VUID-VkCopyImageInfo2-dstImage-01788"></a>
  VUID-VkCopyImageInfo2-dstImage-01788  
  If `dstImage` is of type `VK_IMAGE_TYPE_2D`, then for each element of
  `pRegions`, `dstOffset.z` **must** be `0`

- <a href="#VUID-VkCopyImageInfo2-apiVersion-07933"
  id="VUID-VkCopyImageInfo2-apiVersion-07933"></a>
  VUID-VkCopyImageInfo2-apiVersion-07933  
  If the [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html) extension is
  not enabled, and
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, `srcImage` and `dstImage` **must** have the
  same [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html)

- <a href="#VUID-VkCopyImageInfo2-apiVersion-08969"
  id="VUID-VkCopyImageInfo2-apiVersion-08969"></a>
  VUID-VkCopyImageInfo2-apiVersion-08969  
  If the [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html) extension is
  not enabled, and
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, `srcImage` or `dstImage` is of type
  `VK_IMAGE_TYPE_2D`, then for each element of `pRegions`,
  `extent.depth` **must** be `1`

- <a href="#VUID-VkCopyImageInfo2-srcImage-07743"
  id="VUID-VkCopyImageInfo2-srcImage-07743"></a>
  VUID-VkCopyImageInfo2-srcImage-07743  
  If `srcImage` and `dstImage` have a different
  [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html), and
  [`maintenance5`](#features-maintenance5) is not enabled, one **must**
  be `VK_IMAGE_TYPE_3D` and the other **must** be `VK_IMAGE_TYPE_2D`

- <a href="#VUID-VkCopyImageInfo2-srcImage-08793"
  id="VUID-VkCopyImageInfo2-srcImage-08793"></a>
  VUID-VkCopyImageInfo2-srcImage-08793  
  If `srcImage` and `dstImage` have the same
  [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html), for each element of `pRegions`, if
  neither of the `layerCount` members of `srcSubresource` or
  `dstSubresource` are `VK_REMAINING_ARRAY_LAYERS`, the `layerCount`
  members of `srcSubresource` or `dstSubresource` **must** match

- <a href="#VUID-VkCopyImageInfo2-srcImage-08794"
  id="VUID-VkCopyImageInfo2-srcImage-08794"></a>
  VUID-VkCopyImageInfo2-srcImage-08794  
  If `srcImage` and `dstImage` have the same
  [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html), and one of the `layerCount` members
  of `srcSubresource` or `dstSubresource` is
  `VK_REMAINING_ARRAY_LAYERS`, the other member **must** be either
  `VK_REMAINING_ARRAY_LAYERS` or equal to the `arrayLayers` member of
  the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) used to create the
  image minus `baseArrayLayer`

- <a href="#VUID-VkCopyImageInfo2-srcImage-01790"
  id="VUID-VkCopyImageInfo2-srcImage-01790"></a>
  VUID-VkCopyImageInfo2-srcImage-01790  
  If `srcImage` and `dstImage` are both of type `VK_IMAGE_TYPE_2D`, then
  for each element of `pRegions`, `extent.depth` **must** be `1`

- <a href="#VUID-VkCopyImageInfo2-srcImage-01791"
  id="VUID-VkCopyImageInfo2-srcImage-01791"></a>
  VUID-VkCopyImageInfo2-srcImage-01791  
  If `srcImage` is of type `VK_IMAGE_TYPE_2D`, and `dstImage` is of type
  `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`,
  `extent.depth` **must** equal `srcSubresource.layerCount`

- <a href="#VUID-VkCopyImageInfo2-dstImage-01792"
  id="VUID-VkCopyImageInfo2-dstImage-01792"></a>
  VUID-VkCopyImageInfo2-dstImage-01792  
  If `dstImage` is of type `VK_IMAGE_TYPE_2D`, and `srcImage` is of type
  `VK_IMAGE_TYPE_3D`, then for each element of `pRegions`,
  `extent.depth` **must** equal `dstSubresource.layerCount`

- <a href="#VUID-VkCopyImageInfo2-dstOffset-00150"
  id="VUID-VkCopyImageInfo2-dstOffset-00150"></a>
  VUID-VkCopyImageInfo2-dstOffset-00150  
  For each element of `pRegions`, `dstOffset.x` and (`extent.width` +
  `dstOffset.x`) **must** both be greater than or equal to `0` and less
  than or equal to the width of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-VkCopyImageInfo2-dstOffset-00151"
  id="VUID-VkCopyImageInfo2-dstOffset-00151"></a>
  VUID-VkCopyImageInfo2-dstOffset-00151  
  For each element of `pRegions`, `dstOffset.y` and (`extent.height` +
  `dstOffset.y`) **must** both be greater than or equal to `0` and less
  than or equal to the height of the specified `dstSubresource` of
  `dstImage`

- <a href="#VUID-VkCopyImageInfo2-dstImage-00152"
  id="VUID-VkCopyImageInfo2-dstImage-00152"></a>
  VUID-VkCopyImageInfo2-dstImage-00152  
  If `dstImage` is of type `VK_IMAGE_TYPE_1D`, then for each element of
  `pRegions`, `dstOffset.y` **must** be `0` and `extent.height` **must**
  be `1`

- <a href="#VUID-VkCopyImageInfo2-dstOffset-00153"
  id="VUID-VkCopyImageInfo2-dstOffset-00153"></a>
  VUID-VkCopyImageInfo2-dstOffset-00153  
  If `dstImage` is of type `VK_IMAGE_TYPE_3D`, then for each element of
  `pRegions`, `dstOffset.z` and (`extent.depth` + `dstOffset.z`)
  **must** both be greater than or equal to `0` and less than or equal
  to the depth of the specified `dstSubresource` of `dstImage`

- <a href="#VUID-VkCopyImageInfo2-pRegions-07278"
  id="VUID-VkCopyImageInfo2-pRegions-07278"></a>
  VUID-VkCopyImageInfo2-pRegions-07278  
  For each element of `pRegions`, `srcOffset.x` **must** be a multiple
  of the [texel block extent width](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageInfo2-pRegions-07279"
  id="VUID-VkCopyImageInfo2-pRegions-07279"></a>
  VUID-VkCopyImageInfo2-pRegions-07279  
  For each element of `pRegions`, `srcOffset.y` **must** be a multiple
  of the [texel block extent height](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageInfo2-pRegions-07280"
  id="VUID-VkCopyImageInfo2-pRegions-07280"></a>
  VUID-VkCopyImageInfo2-pRegions-07280  
  For each element of `pRegions`, `srcOffset.z` **must** be a multiple
  of the [texel block extent depth](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageInfo2-pRegions-07281"
  id="VUID-VkCopyImageInfo2-pRegions-07281"></a>
  VUID-VkCopyImageInfo2-pRegions-07281  
  For each element of `pRegions`, `dstOffset.x` **must** be a multiple
  of the [texel block extent width](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageInfo2-pRegions-07282"
  id="VUID-VkCopyImageInfo2-pRegions-07282"></a>
  VUID-VkCopyImageInfo2-pRegions-07282  
  For each element of `pRegions`, `dstOffset.y` **must** be a multiple
  of the [texel block extent height](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageInfo2-pRegions-07283"
  id="VUID-VkCopyImageInfo2-pRegions-07283"></a>
  VUID-VkCopyImageInfo2-pRegions-07283  
  For each element of `pRegions`, `dstOffset.z` **must** be a multiple
  of the [texel block extent depth](#formats-compatibility-classes) of
  the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageInfo2-srcImage-01728"
  id="VUID-VkCopyImageInfo2-srcImage-01728"></a>
  VUID-VkCopyImageInfo2-srcImage-01728  
  For each element of `pRegions`, if the sum of `srcOffset.x` and
  `extent.width` does not equal the width of the subresource specified
  by `srcSubresource`, `extent.width` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageInfo2-srcImage-01729"
  id="VUID-VkCopyImageInfo2-srcImage-01729"></a>
  VUID-VkCopyImageInfo2-srcImage-01729  
  For each element of `pRegions`, if the sum of `srcOffset.y` and
  `extent.height` does not equal the height of the subresource specified
  by `srcSubresource`, `extent.height` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageInfo2-srcImage-01730"
  id="VUID-VkCopyImageInfo2-srcImage-01730"></a>
  VUID-VkCopyImageInfo2-srcImage-01730  
  For each element of `pRegions`, if the sum of `srcOffset.z` and
  `extent.depth` does not equal the depth of the subresource specified
  by `srcSubresource`, `extent.depth` **must** be a multiple of the
  [texel block extent depth](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `srcImage`

- <a href="#VUID-VkCopyImageInfo2-dstImage-01732"
  id="VUID-VkCopyImageInfo2-dstImage-01732"></a>
  VUID-VkCopyImageInfo2-dstImage-01732  
  For each element of `pRegions`, if the sum of `dstOffset.x` and
  `extent.width` does not equal the width of the subresource specified
  by `dstSubresource`, `extent.width` **must** be a multiple of the
  [texel block extent width](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageInfo2-dstImage-01733"
  id="VUID-VkCopyImageInfo2-dstImage-01733"></a>
  VUID-VkCopyImageInfo2-dstImage-01733  
  For each element of `pRegions`, if the sum of `dstOffset.y` and
  `extent.height` does not equal the height of the subresource specified
  by `dstSubresource`, `extent.height` **must** be a multiple of the
  [texel block extent height](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageInfo2-dstImage-01734"
  id="VUID-VkCopyImageInfo2-dstImage-01734"></a>
  VUID-VkCopyImageInfo2-dstImage-01734  
  For each element of `pRegions`, if the sum of `dstOffset.z` and
  `extent.depth` does not equal the depth of the subresource specified
  by `dstSubresource`, `extent.depth` **must** be a multiple of the
  [texel block extent depth](#formats-compatibility-classes) of the
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of `dstImage`

- <a href="#VUID-VkCopyImageInfo2-aspect-06662"
  id="VUID-VkCopyImageInfo2-aspect-06662"></a>
  VUID-VkCopyImageInfo2-aspect-06662  
  If the `aspect` member of any element of `pRegions` includes any flag
  other than `VK_IMAGE_ASPECT_STENCIL_BIT` or `srcImage` was not created
  with [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` **must** have been included in the
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` used to create
  `srcImage`

- <a href="#VUID-VkCopyImageInfo2-aspect-06663"
  id="VUID-VkCopyImageInfo2-aspect-06663"></a>
  VUID-VkCopyImageInfo2-aspect-06663  
  If the `aspect` member of any element of `pRegions` includes any flag
  other than `VK_IMAGE_ASPECT_STENCIL_BIT` or `dstImage` was not created
  with [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` used to create
  `dstImage`

- <a href="#VUID-VkCopyImageInfo2-aspect-06664"
  id="VUID-VkCopyImageInfo2-aspect-06664"></a>
  VUID-VkCopyImageInfo2-aspect-06664  
  If the `aspect` member of any element of `pRegions` includes
  `VK_IMAGE_ASPECT_STENCIL_BIT`, and `srcImage` was created with
  [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` **must** have been included in the
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`
  used to create `srcImage`

- <a href="#VUID-VkCopyImageInfo2-aspect-06665"
  id="VUID-VkCopyImageInfo2-aspect-06665"></a>
  VUID-VkCopyImageInfo2-aspect-06665  
  If the `aspect` member of any element of `pRegions` includes
  `VK_IMAGE_ASPECT_STENCIL_BIT`, and `dstImage` was created with
  [separate stencil usage](#VkImageStencilUsageCreateInfo),
  `VK_IMAGE_USAGE_TRANSFER_DST_BIT` **must** have been included in the
  [VkImageStencilUsageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageStencilUsageCreateInfo.html)::`stencilUsage`
  used to create `dstImage`

<!-- -->

- <a href="#VUID-VkCopyImageInfo2-srcImage-07966"
  id="VUID-VkCopyImageInfo2-srcImage-07966"></a>
  VUID-VkCopyImageInfo2-srcImage-07966  
  If `srcImage` is non-sparse then the image or the specified *disjoint*
  plane **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-VkCopyImageInfo2-srcSubresource-07967"
  id="VUID-VkCopyImageInfo2-srcSubresource-07967"></a>
  VUID-VkCopyImageInfo2-srcSubresource-07967  
  The `srcSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `srcImage` was
  created

- <a href="#VUID-VkCopyImageInfo2-srcSubresource-07968"
  id="VUID-VkCopyImageInfo2-srcSubresource-07968"></a>
  VUID-VkCopyImageInfo2-srcSubresource-07968  
  If `srcSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `srcSubresource.baseArrayLayer` + `srcSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `srcImage` was created

- <a href="#VUID-VkCopyImageInfo2-srcImage-07969"
  id="VUID-VkCopyImageInfo2-srcImage-07969"></a>
  VUID-VkCopyImageInfo2-srcImage-07969  
  `srcImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

<!-- -->

- <a href="#VUID-VkCopyImageInfo2-dstImage-07966"
  id="VUID-VkCopyImageInfo2-dstImage-07966"></a>
  VUID-VkCopyImageInfo2-dstImage-07966  
  If `dstImage` is non-sparse then the image or the specified *disjoint*
  plane **must** be bound completely and contiguously to a single
  `VkDeviceMemory` object

- <a href="#VUID-VkCopyImageInfo2-dstSubresource-07967"
  id="VUID-VkCopyImageInfo2-dstSubresource-07967"></a>
  VUID-VkCopyImageInfo2-dstSubresource-07967  
  The `dstSubresource.mipLevel` member of each element of `pRegions`
  **must** be less than the `mipLevels` specified in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) when `dstImage` was
  created

- <a href="#VUID-VkCopyImageInfo2-dstSubresource-07968"
  id="VUID-VkCopyImageInfo2-dstSubresource-07968"></a>
  VUID-VkCopyImageInfo2-dstSubresource-07968  
  If `dstSubresource.layerCount` is not `VK_REMAINING_ARRAY_LAYERS`,
  `dstSubresource.baseArrayLayer` + `dstSubresource.layerCount` of each
  element of `pRegions` **must** be less than or equal to the
  `arrayLayers` specified in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
  when `dstImage` was created

- <a href="#VUID-VkCopyImageInfo2-dstImage-07969"
  id="VUID-VkCopyImageInfo2-dstImage-07969"></a>
  VUID-VkCopyImageInfo2-dstImage-07969  
  `dstImage` **must** not have been created with `flags` containing
  `VK_IMAGE_CREATE_SUBSAMPLED_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkCopyImageInfo2-sType-sType"
  id="VUID-VkCopyImageInfo2-sType-sType"></a>
  VUID-VkCopyImageInfo2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_IMAGE_INFO_2`

- <a href="#VUID-VkCopyImageInfo2-pNext-pNext"
  id="VUID-VkCopyImageInfo2-pNext-pNext"></a>
  VUID-VkCopyImageInfo2-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkCopyImageInfo2-srcImage-parameter"
  id="VUID-VkCopyImageInfo2-srcImage-parameter"></a>
  VUID-VkCopyImageInfo2-srcImage-parameter  
  `srcImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkCopyImageInfo2-srcImageLayout-parameter"
  id="VUID-VkCopyImageInfo2-srcImageLayout-parameter"></a>
  VUID-VkCopyImageInfo2-srcImageLayout-parameter  
  `srcImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkCopyImageInfo2-dstImage-parameter"
  id="VUID-VkCopyImageInfo2-dstImage-parameter"></a>
  VUID-VkCopyImageInfo2-dstImage-parameter  
  `dstImage` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkCopyImageInfo2-dstImageLayout-parameter"
  id="VUID-VkCopyImageInfo2-dstImageLayout-parameter"></a>
  VUID-VkCopyImageInfo2-dstImageLayout-parameter  
  `dstImageLayout` **must** be a valid
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) value

- <a href="#VUID-VkCopyImageInfo2-pRegions-parameter"
  id="VUID-VkCopyImageInfo2-pRegions-parameter"></a>
  VUID-VkCopyImageInfo2-pRegions-parameter  
  `pRegions` **must** be a valid pointer to an array of `regionCount`
  valid [VkImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy2.html) structures

- <a href="#VUID-VkCopyImageInfo2-regionCount-arraylength"
  id="VUID-VkCopyImageInfo2-regionCount-arraylength"></a>
  VUID-VkCopyImageInfo2-regionCount-arraylength  
  `regionCount` **must** be greater than `0`

- <a href="#VUID-VkCopyImageInfo2-commonparent"
  id="VUID-VkCopyImageInfo2-commonparent"></a>
  VUID-VkCopyImageInfo2-commonparent  
  Both of `dstImage`, and `srcImage` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCopy2.html), [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdCopyImage2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImage2.html),
[vkCmdCopyImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImage2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyImageInfo2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
