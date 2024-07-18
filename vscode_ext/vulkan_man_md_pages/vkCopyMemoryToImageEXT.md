# vkCopyMemoryToImageEXT(3) Manual Page

## Name

vkCopyMemoryToImageEXT - Copy data from host memory into an image



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy data from host memory to an image object, call:

``` c
// Provided by VK_EXT_host_image_copy
VkResult vkCopyMemoryToImageEXT(
    VkDevice                                    device,
    const VkCopyMemoryToImageInfoEXT*           pCopyMemoryToImageInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `pCopyMemoryToImageInfo->dstImage`.

- `pCopyMemoryToImageInfo` is a pointer to a
  [VkCopyMemoryToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageInfoEXT.html)
  structure describing the copy parameters.

## <a href="#_description" class="anchor"></a>Description

This command is functionally similar to
[vkCmdCopyBufferToImage2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage2.html), except it is
executed on the host and reads from host memory instead of a buffer. The
memory of `pCopyMemoryToImageInfo->dstImage` is accessed by the host as
if <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-coherent"
target="_blank" rel="noopener">coherent</a>.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Because queue submissions <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-submission-host-writes"
target="_blank" rel="noopener">automatically make host memory visible to
the device</a>, there would not be a need for a memory barrier before
using the results of this copy operation on the device.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkCopyMemoryToImageEXT-hostImageCopy-09058"
  id="VUID-vkCopyMemoryToImageEXT-hostImageCopy-09058"></a>
  VUID-vkCopyMemoryToImageEXT-hostImageCopy-09058  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-hostImageCopy"
  target="_blank" rel="noopener"><code>hostImageCopy</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCopyMemoryToImageEXT-device-parameter"
  id="VUID-vkCopyMemoryToImageEXT-device-parameter"></a>
  VUID-vkCopyMemoryToImageEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCopyMemoryToImageEXT-pCopyMemoryToImageInfo-parameter"
  id="VUID-vkCopyMemoryToImageEXT-pCopyMemoryToImageInfo-parameter"></a>
  VUID-vkCopyMemoryToImageEXT-pCopyMemoryToImageInfo-parameter  
  `pCopyMemoryToImageInfo` **must** be a valid pointer to a valid
  [VkCopyMemoryToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageInfoEXT.html)
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
[VkCopyMemoryToImageInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageInfoEXT.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCopyMemoryToImageEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
