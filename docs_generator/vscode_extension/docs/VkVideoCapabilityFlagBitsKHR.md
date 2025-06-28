# VkVideoCapabilityFlagBitsKHR(3) Manual Page

## Name

VkVideoCapabilityFlagBitsKHR - Video decode and encode capability bits



## [](#_c_specification)C Specification

Bits which **can** be set in [VkVideoCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilitiesKHR.html)::`flags` are:

```c++
// Provided by VK_KHR_video_queue
typedef enum VkVideoCapabilityFlagBitsKHR {
    VK_VIDEO_CAPABILITY_PROTECTED_CONTENT_BIT_KHR = 0x00000001,
    VK_VIDEO_CAPABILITY_SEPARATE_REFERENCE_IMAGES_BIT_KHR = 0x00000002,
} VkVideoCapabilityFlagBitsKHR;
```

## [](#_description)Description

- `VK_VIDEO_CAPABILITY_PROTECTED_CONTENT_BIT_KHR` specifies that video sessions support producing and consuming protected content.
- []()`VK_VIDEO_CAPABILITY_SEPARATE_REFERENCE_IMAGES_BIT_KHR` indicates that the [video picture resources](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#video-picture-resources) associated with the [DPB slots](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#dpb-slot) of a video session **can** be backed by separate `VkImage` objects. If this capability flag is not present, then all DPB slots of a video session **must** be associated with video picture resources backed by the same `VkImage` object (e.g. using different layers of the same image).

## [](#_see_also)See Also

[VK\_KHR\_video\_queue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_video_queue.html), [VkVideoCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoCapabilityFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkVideoCapabilityFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0