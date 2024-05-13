# vkCopyImageToImageEXT(3) Manual Page

## Name

vkCopyImageToImageEXT - Copy image data using the host



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data from an image object to another image object using the
host, call:

``` c
// Provided by VK_EXT_host_image_copy
VkResult vkCopyImageToImageEXT(
    VkDevice                                    device,
    const VkCopyImageToImageInfoEXT*            pCopyImageToImageInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `pCopyImageToMemoryInfo->srcImage`.

- `pCopyImageToImageInfo` is a pointer to a
  [VkCopyImageToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToImageInfoEXT.html) structure
  describing the copy parameters.

## <a href="#_description" class="anchor"></a>Description

This command is functionally similar to
[vkCmdCopyImage2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImage2.html), except it is executed on the
host.

Valid Usage

- <a href="#VUID-vkCopyImageToImageEXT-hostImageCopy-09068"
  id="VUID-vkCopyImageToImageEXT-hostImageCopy-09068"></a>
  VUID-vkCopyImageToImageEXT-hostImageCopy-09068  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-hostImageCopy"
  target="_blank" rel="noopener"><code>hostImageCopy</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCopyImageToImageEXT-device-parameter"
  id="VUID-vkCopyImageToImageEXT-device-parameter"></a>
  VUID-vkCopyImageToImageEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCopyImageToImageEXT-pCopyImageToImageInfo-parameter"
  id="VUID-vkCopyImageToImageEXT-pCopyImageToImageInfo-parameter"></a>
  VUID-vkCopyImageToImageEXT-pCopyImageToImageInfo-parameter  
  `pCopyImageToImageInfo` **must** be a valid pointer to a valid
  [VkCopyImageToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToImageInfoEXT.html) structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_INITIALIZATION_FAILED`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_MEMORY_MAP_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_host_image_copy](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_host_image_copy.html),
[VkCopyImageToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToImageInfoEXT.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCopyImageToImageEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
