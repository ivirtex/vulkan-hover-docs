# VkPipelineRobustnessImageBehaviorEXT(3) Manual Page

## Name

VkPipelineRobustnessImageBehaviorEXT - Enum controlling the robustness
of image accesses in a pipeline stage



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the `images` member of
[VkPipelineRobustnessCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessCreateInfoEXT.html)
are:

``` c
// Provided by VK_EXT_pipeline_robustness
typedef enum VkPipelineRobustnessImageBehaviorEXT {
    VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_DEVICE_DEFAULT_EXT = 0,
    VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_DISABLED_EXT = 1,
    VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS_EXT = 2,
    VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS_2_EXT = 3,
} VkPipelineRobustnessImageBehaviorEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_DEVICE_DEFAULT_EXT` specifies
  that this pipeline stage follows the behavior of robustness features
  that are enabled on the device that created this pipeline

- `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_DISABLED_EXT` specifies that
  image accesses by this pipeline stage to the relevant resource types
  **must** not be out of bounds

- `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS_EXT`
  specifies that out of bounds accesses by this pipeline stage to images
  behave as if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustImageAccess"
  target="_blank" rel="noopener"><code>robustImageAccess</code></a>
  feature is enabled

- `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS_2_EXT`
  specifies that out of bounds accesses by this pipeline stage to images
  behave as if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustImageAccess2"
  target="_blank" rel="noopener"><code>robustImageAccess2</code></a>
  feature is enabled

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_pipeline_robustness](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_pipeline_robustness.html),
[VkPhysicalDevicePipelineRobustnessPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineRobustnessPropertiesEXT.html),
[VkPipelineRobustnessCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessCreateInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineRobustnessImageBehaviorEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
