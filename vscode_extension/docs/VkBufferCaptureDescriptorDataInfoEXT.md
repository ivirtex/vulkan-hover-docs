# VkBufferCaptureDescriptorDataInfoEXT(3) Manual Page

## Name

VkBufferCaptureDescriptorDataInfoEXT - Structure specifying a buffer for
descriptor capture



## <a href="#_c_specification" class="anchor"></a>C Specification

Information about the buffer to get descriptor buffer capture data for
is passed in a `VkBufferCaptureDescriptorDataInfoEXT` structure:

``` c
// Provided by VK_EXT_descriptor_buffer
typedef struct VkBufferCaptureDescriptorDataInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBuffer           buffer;
} VkBufferCaptureDescriptorDataInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `buffer` is the `VkBuffer` handle of the buffer to get opaque capture
  data for.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBufferCaptureDescriptorDataInfoEXT-buffer-08075"
  id="VUID-VkBufferCaptureDescriptorDataInfoEXT-buffer-08075"></a>
  VUID-VkBufferCaptureDescriptorDataInfoEXT-buffer-08075  
  `buffer` **must** have been created with
  `VK_BUFFER_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` set in
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html)::`flags`

Valid Usage (Implicit)

- <a href="#VUID-VkBufferCaptureDescriptorDataInfoEXT-sType-sType"
  id="VUID-VkBufferCaptureDescriptorDataInfoEXT-sType-sType"></a>
  VUID-VkBufferCaptureDescriptorDataInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BUFFER_CAPTURE_DESCRIPTOR_DATA_INFO_EXT`

- <a href="#VUID-VkBufferCaptureDescriptorDataInfoEXT-pNext-pNext"
  id="VUID-VkBufferCaptureDescriptorDataInfoEXT-pNext-pNext"></a>
  VUID-VkBufferCaptureDescriptorDataInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkBufferCaptureDescriptorDataInfoEXT-buffer-parameter"
  id="VUID-VkBufferCaptureDescriptorDataInfoEXT-buffer-parameter"></a>
  VUID-VkBufferCaptureDescriptorDataInfoEXT-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetBufferOpaqueCaptureDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferOpaqueCaptureDescriptorDataEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferCaptureDescriptorDataInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
