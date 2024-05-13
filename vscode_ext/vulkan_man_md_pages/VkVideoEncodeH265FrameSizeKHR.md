# VkVideoEncodeH265FrameSizeKHR(3) Manual Page

## Name

VkVideoEncodeH265FrameSizeKHR - Structure describing frame size values
per H.265 picture type



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkVideoEncodeH265FrameSizeKHR` structure is defined as:

``` c
// Provided by VK_KHR_video_encode_h265
typedef struct VkVideoEncodeH265FrameSizeKHR {
    uint32_t    frameISize;
    uint32_t    framePSize;
    uint32_t    frameBSize;
} VkVideoEncodeH265FrameSizeKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `frameISize` is the size in bytes to be used for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-i-pic"
  target="_blank" rel="noopener">I frames</a>.

- `framePSize` is the size in bytes to be used for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-p-pic"
  target="_blank" rel="noopener">P frames</a>.

- `frameBSize` is the size in bytes to be used for <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#encode-h265-b-pic"
  target="_blank" rel="noopener">B frames</a>.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_encode_h265](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_encode_h265.html),
[VkVideoEncodeH265RateControlLayerInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlLayerInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoEncodeH265FrameSizeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
