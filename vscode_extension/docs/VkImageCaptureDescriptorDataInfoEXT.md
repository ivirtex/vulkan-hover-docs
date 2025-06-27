# VkImageCaptureDescriptorDataInfoEXT(3) Manual Page

## Name

VkImageCaptureDescriptorDataInfoEXT - Structure specifying an image for
descriptor capture



## <a href="#_c_specification" class="anchor"></a>C Specification

Information about the image to get descriptor buffer capture data for is
passed in a `VkImageCaptureDescriptorDataInfoEXT` structure:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkImageCaptureDescriptorDataInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
} VkImageCaptureDescriptorDataInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `image` is the `VkImage` handle of the image to get opaque capture
  data for.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImageCaptureDescriptorDataInfoEXT-image-08079"
  id="VUID-VkImageCaptureDescriptorDataInfoEXT-image-08079"></a>
  VUID-VkImageCaptureDescriptorDataInfoEXT-image-08079  
  `image` **must** have been created with
  `VK_IMAGE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` set in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags`

Valid Usage (Implicit)

- <a href="#VUID-VkImageCaptureDescriptorDataInfoEXT-sType-sType"
  id="VUID-VkImageCaptureDescriptorDataInfoEXT-sType-sType"></a>
  VUID-VkImageCaptureDescriptorDataInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`

- <a href="#VUID-VkImageCaptureDescriptorDataInfoEXT-pNext-pNext"
  id="VUID-VkImageCaptureDescriptorDataInfoEXT-pNext-pNext"></a>
  VUID-VkImageCaptureDescriptorDataInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkImageCaptureDescriptorDataInfoEXT-image-parameter"
  id="VUID-VkImageCaptureDescriptorDataInfoEXT-image-parameter"></a>
  VUID-VkImageCaptureDescriptorDataInfoEXT-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetImageOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageOpaqueCaptureDescriptorDataEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageCaptureDescriptorDataInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
