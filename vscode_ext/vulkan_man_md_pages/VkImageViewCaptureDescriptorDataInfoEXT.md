# VkImageViewCaptureDescriptorDataInfoEXT(3) Manual Page

## Name

VkImageViewCaptureDescriptorDataInfoEXT - Structure specifying an image
view for descriptor capture



## <a href="#_c_specification" class="anchor"></a>C Specification

Information about the image view to get descriptor buffer capture data
for is passed in a `VkImageViewCaptureDescriptorDataInfoEXT` structure:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkImageViewCaptureDescriptorDataInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkImageView        imageView;
} VkImageViewCaptureDescriptorDataInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `imageView` is the `VkImageView` handle of the image view to get
  opaque capture data for.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImageViewCaptureDescriptorDataInfoEXT-imageView-08083"
  id="VUID-VkImageViewCaptureDescriptorDataInfoEXT-imageView-08083"></a>
  VUID-VkImageViewCaptureDescriptorDataInfoEXT-imageView-08083  
  `imageView` **must** have been created with
  `VK_IMAGE_VIEW_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` set in
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)::`flags`

Valid Usage (Implicit)

- <a href="#VUID-VkImageViewCaptureDescriptorDataInfoEXT-sType-sType"
  id="VUID-VkImageViewCaptureDescriptorDataInfoEXT-sType-sType"></a>
  VUID-VkImageViewCaptureDescriptorDataInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_VIEW_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`

- <a href="#VUID-VkImageViewCaptureDescriptorDataInfoEXT-pNext-pNext"
  id="VUID-VkImageViewCaptureDescriptorDataInfoEXT-pNext-pNext"></a>
  VUID-VkImageViewCaptureDescriptorDataInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkImageViewCaptureDescriptorDataInfoEXT-imageView-parameter"
  id="VUID-VkImageViewCaptureDescriptorDataInfoEXT-imageView-parameter"></a>
  VUID-VkImageViewCaptureDescriptorDataInfoEXT-imageView-parameter  
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetImageViewOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageViewOpaqueCaptureDescriptorDataEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageViewCaptureDescriptorDataInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
