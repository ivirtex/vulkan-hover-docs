# VkVideoSessionCreateFlagBitsKHR(3) Manual Page

## Name

VkVideoSessionCreateFlagBitsKHR - Video session creation flags



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkVideoSessionCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateInfoKHR.html)::`flags`
are:

``` c
// Provided by VK_KHR_video_queue
typedef enum VkVideoSessionCreateFlagBitsKHR {
    VK_VIDEO_SESSION_CREATE_PROTECTED_CONTENT_BIT_KHR = 0x00000001,
  // Provided by VK_KHR_video_encode_queue
    VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_PARAMETER_OPTIMIZATIONS_BIT_KHR = 0x00000002,
  // Provided by VK_KHR_video_maintenance1
    VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR = 0x00000004,
} VkVideoSessionCreateFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_SESSION_CREATE_PROTECTED_CONTENT_BIT_KHR` specifies that the
  video session uses protected video content.

- <span id="encode-optimizing-overrides"></span>
  `VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_PARAMETER_OPTIMIZATIONS_BIT_KHR`
  specifies that the implementation is allowed to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-overrides"
  target="_blank" rel="noopener">override</a> video session parameters
  and other codec-specific encoding parameters to optimize video encode
  operations based on the use case information specified in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-profiles"
  target="_blank" rel="noopener">video profile</a> and the used <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-quality-level"
  target="_blank" rel="noopener">video encode quality level</a>.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>Not specifying
  <code>VK_VIDEO_SESSION_CREATE_ALLOW_ENCODE_PARAMETER_OPTIMIZATIONS_BIT_KHR</code>
  does not guarantee that the implementation will not do any
  codec-specific parameter overrides, as certain overrides are necessary
  for the correct operation of the video encoder implementation due to
  limitations to the available encoding tools on that implementation. This
  flag, however, enables the implementation to apply further optimizing
  overrides.</p></td>
  </tr>
  </tbody>
  </table>

- `VK_VIDEO_SESSION_CREATE_INLINE_QUERIES_BIT_KHR` specifies that
  queries within video coding scopes using the created video session are
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-inline-queries"
  target="_blank" rel="noopener">executed inline</a> with video coding
  operations.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkVideoSessionCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoSessionCreateFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
