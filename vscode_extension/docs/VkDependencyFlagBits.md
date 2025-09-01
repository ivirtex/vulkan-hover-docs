# VkDependencyFlagBits(3) Manual Page

## Name

VkDependencyFlagBits - Bitmask specifying how execution and memory dependencies are formed



## [](#_c_specification)C Specification

Bits which **can** be set in `vkCmdPipelineBarrier`::`dependencyFlags`, specifying how execution and memory dependencies are formed, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkDependencyFlagBits {
    VK_DEPENDENCY_BY_REGION_BIT = 0x00000001,
  // Provided by VK_VERSION_1_1
    VK_DEPENDENCY_DEVICE_GROUP_BIT = 0x00000004,
  // Provided by VK_VERSION_1_1
    VK_DEPENDENCY_VIEW_LOCAL_BIT = 0x00000002,
  // Provided by VK_EXT_attachment_feedback_loop_layout
    VK_DEPENDENCY_FEEDBACK_LOOP_BIT_EXT = 0x00000008,
  // Provided by VK_KHR_maintenance8
    VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR = 0x00000020,
  // Provided by VK_KHR_maintenance9
    VK_DEPENDENCY_ASYMMETRIC_EVENT_BIT_KHR = 0x00000040,
  // Provided by VK_KHR_multiview
    VK_DEPENDENCY_VIEW_LOCAL_BIT_KHR = VK_DEPENDENCY_VIEW_LOCAL_BIT,
  // Provided by VK_KHR_device_group
    VK_DEPENDENCY_DEVICE_GROUP_BIT_KHR = VK_DEPENDENCY_DEVICE_GROUP_BIT,
} VkDependencyFlagBits;
```

## [](#_description)Description

- `VK_DEPENDENCY_BY_REGION_BIT` specifies that dependencies will be [framebuffer-local](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-framebuffer-regions).
- `VK_DEPENDENCY_VIEW_LOCAL_BIT` specifies that dependencies will be [view-local](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-view-local-dependencies).
- `VK_DEPENDENCY_DEVICE_GROUP_BIT` specifies that dependencies are [non-device-local](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-device-local-dependencies).
- `VK_DEPENDENCY_FEEDBACK_LOOP_BIT_EXT` specifies that the render pass will write to and read from the same image with [feedback loop enabled](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-feedbackloop).
- `VK_DEPENDENCY_QUEUE_FAMILY_OWNERSHIP_TRANSFER_USE_ALL_STAGES_BIT_KHR` specifies that source and destination stages are not ignored when performing a [queue family ownership transfer](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-queue-transfers).

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDependencyFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDependencyFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDependencyFlagBits).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0