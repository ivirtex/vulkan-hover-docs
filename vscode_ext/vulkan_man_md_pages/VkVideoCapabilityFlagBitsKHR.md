# VkVideoCapabilityFlagBitsKHR(3) Manual Page

## Name

VkVideoCapabilityFlagBitsKHR - Video decode and encode capability bits



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilitiesKHR.html)::`flags` are:

``` c
// Provided by VK_KHR_video_queue
typedef enum VkVideoCapabilityFlagBitsKHR {
    VK_VIDEO_CAPABILITY_PROTECTED_CONTENT_BIT_KHR = 0x00000001,
    VK_VIDEO_CAPABILITY_SEPARATE_REFERENCE_IMAGES_BIT_KHR = 0x00000002,
} VkVideoCapabilityFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_VIDEO_CAPABILITY_PROTECTED_CONTENT_BIT_KHR` indicates that video
  sessions support producing and consuming protected content.

- <span id="separate-reference-images"></span>
  `VK_VIDEO_CAPABILITY_SEPARATE_REFERENCE_IMAGES_BIT_KHR` indicates that
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#video-picture-resources"
  target="_blank" rel="noopener">video picture resources</a> associated
  with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#dpb-slot"
  target="_blank" rel="noopener">DPB slots</a> of a video session
  **can** be backed by separate `VkImage` objects. If this capability
  flag is not present, then all DPB slots of a video session **must** be
  associated with video picture resources backed by the same `VkImage`
  object (e.g. using different layers of the same image).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_video_queue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_video_queue.html),
[VkVideoCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilityFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkVideoCapabilityFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
