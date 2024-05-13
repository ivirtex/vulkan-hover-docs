# VkShaderFloatControlsIndependence(3) Manual Page

## Name

VkShaderFloatControlsIndependence - Bitmask specifying whether, and how,
shader float controls can be set separately



## <a href="#_c_specification" class="anchor"></a>C Specification

Values which **may** be returned in the `denormBehaviorIndependence` and
`roundingModeIndependence` fields of
`VkPhysicalDeviceFloatControlsProperties` are:

``` c
// Provided by VK_VERSION_1_2
typedef enum VkShaderFloatControlsIndependence {
    VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY = 0,
    VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL = 1,
    VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE = 2,
  // Provided by VK_KHR_shader_float_controls
    VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY_KHR = VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY,
  // Provided by VK_KHR_shader_float_controls
    VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL_KHR = VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL,
  // Provided by VK_KHR_shader_float_controls
    VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE_KHR = VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE,
} VkShaderFloatControlsIndependence;
```

or the equivalent

``` c
// Provided by VK_KHR_shader_float_controls
typedef VkShaderFloatControlsIndependence VkShaderFloatControlsIndependenceKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_32_BIT_ONLY` specifies that
  shader float controls for 32-bit floating point **can** be set
  independently; other bit widths **must** be set identically to each
  other.

- `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_ALL` specifies that shader
  float controls for all bit widths **can** be set independently.

- `VK_SHADER_FLOAT_CONTROLS_INDEPENDENCE_NONE` specifies that shader
  float controls for all bit widths **must** be set identically.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_shader_float_controls](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_float_controls.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkPhysicalDeviceFloatControlsProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFloatControlsProperties.html),
[VkPhysicalDeviceVulkan12Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Properties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShaderFloatControlsIndependence"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
