# vkCopyImageToMemoryEXT(3) Manual Page

## Name

vkCopyImageToMemoryEXT - Copy image data into host memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data from an image object to host memory, call:

``` c
// Provided by VK_EXT_host_image_copy
VkResult vkCopyImageToMemoryEXT(
    VkDevice                                    device,
    const VkCopyImageToMemoryInfoEXT*           pCopyImageToMemoryInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `pCopyImageToMemoryInfo->srcImage`.

- `pCopyImageToMemoryInfo` is a pointer to a
  [VkCopyImageToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToMemoryInfoEXT.html)
  structure describing the copy parameters.

## <a href="#_description" class="anchor"></a>Description

This command is functionally similar to
[vkCmdCopyImageToBuffer2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImageToBuffer2.html), except it is
executed on the host and writes to host memory instead of a buffer.

Valid Usage

- <a href="#VUID-vkCopyImageToMemoryEXT-hostImageCopy-09063"
  id="VUID-vkCopyImageToMemoryEXT-hostImageCopy-09063"></a>
  VUID-vkCopyImageToMemoryEXT-hostImageCopy-09063  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-hostImageCopy"
  target="_blank" rel="noopener"><code>hostImageCopy</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCopyImageToMemoryEXT-device-parameter"
  id="VUID-vkCopyImageToMemoryEXT-device-parameter"></a>
  VUID-vkCopyImageToMemoryEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCopyImageToMemoryEXT-pCopyImageToMemoryInfo-parameter"
  id="VUID-vkCopyImageToMemoryEXT-pCopyImageToMemoryInfo-parameter"></a>
  VUID-vkCopyImageToMemoryEXT-pCopyImageToMemoryInfo-parameter  
  `pCopyImageToMemoryInfo` **must** be a valid pointer to a valid
  [VkCopyImageToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToMemoryInfoEXT.html)
  structure

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
[VkCopyImageToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyImageToMemoryInfoEXT.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCopyImageToMemoryEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
