# VkVideoEncodeContentFlagBitsKHR(3) Manual Page

## Name

VkVideoEncodeContentFlagBitsKHR - Video encode content flags



## <a href="#_c_specification" class="anchor"></a>C Specification

The following bits **can** be specified in
[VkVideoEncodeUsageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeUsageInfoKHR.html)::`videoContentHints`
as a hint about the encoded video content:

``` c
// Provided by VK_KHR_video_encode_queue
typedef enum VkVideoEncodeContentFlagBitsKHR {
    VK_VIDEO_ENCODE_CONTENT_DEFAULT_KHR = 0,
    VK_VIDEO_ENCODE_CONTENT_CAMERA_BIT_KHR = 0x00000001,
    VK_VIDEO_ENCODE_CONTENT_DESKTOP_BIT_KHR = 0x00000002,
    VK_VIDEO_ENCODE_CONTENT_RENDERED_BIT_KHR = 0x00000004,
} VkVideoEncodeContentFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_ENCODE_CONTENT_CAMERA_BIT_KHR` specifies that video encoding
  is intended to be used to encode camera content.

- `VK_VIDEO_ENCODE_CONTENT_DESKTOP_BIT_KHR` specifies that video
  encoding is intended to be used to encode desktop content.

- `VK_VIDEO_ENCODE_CONTENT_RENDERED_BIT_KHR` specified that video
  encoding is intended to be used to encode rendered (e.g. game)
  content.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>There are no restrictions on the combination of bits that
<strong>can</strong> be specified by the application. However,
applications <strong>should</strong> use reasonable combinations in
order for the implementation to be able to select the most appropriate
mode of operation for the particular content type.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_queue.html),
[VkVideoEncodeContentFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeContentFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeContentFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
