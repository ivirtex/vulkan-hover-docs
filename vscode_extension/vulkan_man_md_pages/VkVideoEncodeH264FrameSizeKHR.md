# VkVideoEncodeH264FrameSizeKHR(3) Manual Page

## Name

VkVideoEncodeH264FrameSizeKHR - Structure describing frame size values
per H.264 picture type



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeH264FrameSizeKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h264
typedef struct VkVideoEncodeH264FrameSizeKHR {
    uint32_t    frameISize;
    uint32_t    framePSize;
    uint32_t    frameBSize;
} VkVideoEncodeH264FrameSizeKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `frameISize` is the size in bytes to be used for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-i-pic"
  target="_blank" rel="noopener">I pictures</a>.

- `framePSize` is the size in bytes to be used for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-p-pic"
  target="_blank" rel="noopener">P pictures</a>.

- `frameBSize` is the size in bytes to be used for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h264-b-pic"
  target="_blank" rel="noopener">B pictures</a>.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h264](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h264.html),
[VkVideoEncodeH264RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlLayerInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH264FrameSizeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
