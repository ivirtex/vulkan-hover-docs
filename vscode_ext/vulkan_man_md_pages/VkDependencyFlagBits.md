# VkDependencyFlagBits(3) Manual Page

## Name

VkDependencyFlagBits - Bitmask specifying how execution and memory
dependencies are formed



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in `vkCmdPipelineBarrier`::`dependencyFlags`,
specifying how execution and memory dependencies are formed, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkDependencyFlagBits {
    VK_DEPENDENCY_BY_REGION_BIT = 0x00000001,
  // Provided by VK_VERSION_1_1
    VK_DEPENDENCY_DEVICE_GROUP_BIT = 0x00000004,
  // Provided by VK_VERSION_1_1
    VK_DEPENDENCY_VIEW_LOCAL_BIT = 0x00000002,
  // Provided by VK_EXT_attachment_feedback_loop_layout
    VK_DEPENDENCY_FEEDBACK_LOOP_BIT_EXT = 0x00000008,
  // Provided by VK_KHR_multiview
    VK_DEPENDENCY_VIEW_LOCAL_BIT_KHR = VK_DEPENDENCY_VIEW_LOCAL_BIT,
  // Provided by VK_KHR_device_group
    VK_DEPENDENCY_DEVICE_GROUP_BIT_KHR = VK_DEPENDENCY_DEVICE_GROUP_BIT,
} VkDependencyFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEPENDENCY_BY_REGION_BIT` specifies that dependencies will be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-framebuffer-regions"
  target="_blank" rel="noopener">framebuffer-local</a>.

- `VK_DEPENDENCY_VIEW_LOCAL_BIT` specifies that dependencies will be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-view-local-dependencies"
  target="_blank" rel="noopener">view-local</a>.

- `VK_DEPENDENCY_DEVICE_GROUP_BIT` specifies that dependencies are <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-device-local-dependencies"
  target="_blank" rel="noopener">non-device-local</a>.

- `VK_DEPENDENCY_FEEDBACK_LOOP_BIT_EXT` specifies that the render pass
  will write to and read from the same image using the
  `VK_IMAGE_LAYOUT_ATTACHMENT_FEEDBACK_LOOP_OPTIMAL_EXT` layout.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDependencyFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDependencyFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
