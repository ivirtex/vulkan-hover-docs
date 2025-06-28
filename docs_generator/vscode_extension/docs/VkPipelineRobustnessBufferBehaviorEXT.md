# VkPipelineRobustnessBufferBehavior(3) Manual Page

## Name

VkPipelineRobustnessBufferBehavior - Enum controlling the robustness of buffer accesses in a pipeline stage



## [](#_c_specification)C Specification

Possible values of the `storageBuffers`, `uniformBuffers`, and `vertexInputs` members of [VkPipelineRobustnessCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessCreateInfo.html) are:

```c++
// Provided by VK_VERSION_1_4
typedef enum VkPipelineRobustnessBufferBehavior {
    VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DEVICE_DEFAULT = 0,
    VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DISABLED = 1,
    VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS = 2,
    VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2 = 3,
  // Provided by VK_EXT_pipeline_robustness
    VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DEVICE_DEFAULT_EXT = VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DEVICE_DEFAULT,
  // Provided by VK_EXT_pipeline_robustness
    VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DISABLED_EXT = VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DISABLED,
  // Provided by VK_EXT_pipeline_robustness
    VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_EXT = VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS,
  // Provided by VK_EXT_pipeline_robustness
    VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2_EXT = VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2,
} VkPipelineRobustnessBufferBehavior;
```

or the equivalent

```c++
// Provided by VK_EXT_pipeline_robustness
typedef VkPipelineRobustnessBufferBehavior VkPipelineRobustnessBufferBehaviorEXT;
```

## [](#_description)Description

- `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DEVICE_DEFAULT` specifies that [out of bounds](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-execution-memory-access-bounds) buffer accesses follow the behavior of robust buffer access features enabled for the device.
- `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DISABLED` specifies that buffer accesses **must** not be [out of bounds](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-execution-memory-access-bounds).
- `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS` specifies that buffer accesses conform to [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-robust-buffer-access](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-robust-buffer-access) guarantees.
- `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2` specifies that buffer accesses conform to [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-robust-buffer-access2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#shaders-robust-buffer-access2) guarantees.

## [](#_see_also)See Also

[VK\_EXT\_pipeline\_robustness](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_robustness.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkPhysicalDevicePipelineRobustnessProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineRobustnessProperties.html), [VkPhysicalDeviceVulkan14Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVulkan14Properties.html), [VkPipelineRobustnessCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineRobustnessBufferBehavior)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0