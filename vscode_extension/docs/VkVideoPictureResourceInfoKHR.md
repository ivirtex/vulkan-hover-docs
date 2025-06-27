# VkVideoPictureResourceInfoKHR(3) Manual Page

## Name

VkVideoPictureResourceInfoKHR - Structure specifying the parameters of a
video picture resource



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoPictureResourceInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_queue
typedef struct VkVideoPictureResourceInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkOffset2D         codedOffset;
    VkExtent2D         codedExtent;
    uint32_t           baseArrayLayer;
    VkImageView        imageViewBinding;
} VkVideoPictureResourceInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `codedOffset` is the offset in texels of the image subregion to use.

- `codedExtent` is the size in pixels of the coded image data.

- `baseArrayLayer` is the array layer of the image view specified in
  `imageViewBinding` to use as the video picture resource.

- `imageViewBinding` is an image view representing the video picture
  resource.

## <a href="#_description" class="anchor"></a>Description

The image subresource referred to by such a structure is defined as the
image array layer index specified in `baseArrayLayer` relative to the
image subresource range the image view specified in `imageViewBinding`
was created with.

The meaning of the `codedOffset` and `codedExtent` depends on the
command and context the video picture resource is used in, as well as on
the used <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profiles"
target="_blank" rel="noopener">video profile</a> and corresponding
codec-specific semantics, as described later.

A video picture resource is uniquely defined by the image subresource
referred to by an instance of this structure, together with the
`codedOffset` and `codedExtent` members that identify the image
subregion within the image subresource referenced corresponding to the
video picture resource according to the particular codec-specific
semantics.

Accesses to image data within a video picture resource happen at the
granularity indicated by
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`pictureAccessGranularity`,
as returned by
[vkGetPhysicalDeviceVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceVideoCapabilitiesKHR.html)
for the used <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profiles"
target="_blank" rel="noopener">video profile</a>. As a result, given an
effective image subregion corresponding to a video picture resource, the
actual image subregion accessed **may** be larger than that as it
**may** include additional padding texels due to the picture access
granularity. Any writes performed by video coding operations to such
padding texels will result in undefined texel values.

Two video picture resources match if they refer to the same image
subresource and they specify identical `codedOffset` and `codedExtent`
values.

Valid Usage

- <a href="#VUID-VkVideoPictureResourceInfoKHR-baseArrayLayer-07175"
  id="VUID-VkVideoPictureResourceInfoKHR-baseArrayLayer-07175"></a>
  VUID-VkVideoPictureResourceInfoKHR-baseArrayLayer-07175  
  `baseArrayLayer` **must** be less than the
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)::`subresourceRange.layerCount`
  specified when the image view `imageViewBinding` was created

Valid Usage (Implicit)

- <a href="#VUID-VkVideoPictureResourceInfoKHR-sType-sType"
  id="VUID-VkVideoPictureResourceInfoKHR-sType-sType"></a>
  VUID-VkVideoPictureResourceInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_VIDEO_PICTURE_RESOURCE_INFO_KHR`

- <a href="#VUID-VkVideoPictureResourceInfoKHR-pNext-pNext"
  id="VUID-VkVideoPictureResourceInfoKHR-pNext-pNext"></a>
  VUID-VkVideoPictureResourceInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkVideoPictureResourceInfoKHR-imageViewBinding-parameter"
  id="VUID-VkVideoPictureResourceInfoKHR-imageViewBinding-parameter"></a>
  VUID-VkVideoPictureResourceInfoKHR-imageViewBinding-parameter  
  `imageViewBinding` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)
  handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html),
[VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkVideoDecodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeInfoKHR.html),
[VkVideoEncodeInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeInfoKHR.html),
[VkVideoReferenceSlotInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoReferenceSlotInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoPictureResourceInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
