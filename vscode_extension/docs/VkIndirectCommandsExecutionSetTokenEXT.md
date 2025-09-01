# VkIndirectCommandsExecutionSetTokenEXT(3) Manual Page

## Name

VkIndirectCommandsExecutionSetTokenEXT - Structure specifying input data for a single execution set command token



## [](#_c_specification)C Specification

The `VkIndirectCommandsExecutionSetTokenEXT` structure specifies the input data for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_EXECUTION_SET_EXT` token.

```c++
// Provided by VK_EXT_device_generated_commands
typedef struct VkIndirectCommandsExecutionSetTokenEXT {
    VkIndirectExecutionSetInfoTypeEXT    type;
    VkShaderStageFlags                   shaderStages;
} VkIndirectCommandsExecutionSetTokenEXT;
```

## [](#_members)Members

- `type` describes the type of indirect execution set in use.
- `shaderStages` specifies the shaders that will be changed by this token.

## [](#_description)Description

Valid Usage

- [](#VUID-VkIndirectCommandsExecutionSetTokenEXT-shaderStages-11137)VUID-VkIndirectCommandsExecutionSetTokenEXT-shaderStages-11137  
  Each bit in `shaderStages` **must** be supported by [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-supportedIndirectCommandsShaderStagesPipelineBinding)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`supportedIndirectCommandsShaderStagesPipelineBinding` or [](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-supportedIndirectCommandsShaderStagesShaderBinding)[VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesEXT.html)::`supportedIndirectCommandsShaderStagesShaderBinding`

Valid Usage (Implicit)

- [](#VUID-VkIndirectCommandsExecutionSetTokenEXT-type-parameter)VUID-VkIndirectCommandsExecutionSetTokenEXT-type-parameter  
  `type` **must** be a valid [VkIndirectExecutionSetInfoTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetInfoTypeEXT.html) value
- [](#VUID-VkIndirectCommandsExecutionSetTokenEXT-shaderStages-parameter)VUID-VkIndirectCommandsExecutionSetTokenEXT-shaderStages-parameter  
  `shaderStages` **must** be a valid combination of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) values
- [](#VUID-VkIndirectCommandsExecutionSetTokenEXT-shaderStages-requiredbitmask)VUID-VkIndirectCommandsExecutionSetTokenEXT-shaderStages-requiredbitmask  
  `shaderStages` **must** not be `0`

## [](#_see_also)See Also

[VK\_EXT\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_generated_commands.html), [VkIndirectCommandsTokenDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsTokenDataEXT.html), [VkIndirectExecutionSetInfoTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetInfoTypeEXT.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkIndirectCommandsExecutionSetTokenEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0