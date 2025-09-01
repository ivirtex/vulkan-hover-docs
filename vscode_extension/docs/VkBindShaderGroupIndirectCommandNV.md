# VkBindShaderGroupIndirectCommandNV(3) Manual Page

## Name

VkBindShaderGroupIndirectCommandNV - Structure specifying input data for a single shader group command token



## [](#_c_specification)C Specification

The `VkBindShaderGroupIndirectCommandNV` structure specifies the input data for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SHADER_GROUP_NV` token.

```c++
// Provided by VK_NV_device_generated_commands
typedef struct VkBindShaderGroupIndirectCommandNV {
    uint32_t    groupIndex;
} VkBindShaderGroupIndirectCommandNV;
```

## [](#_members)Members

- `groupIndex` specifies which shader group of the current bound graphics pipeline is used.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBindShaderGroupIndirectCommandNV-None-02944)VUID-VkBindShaderGroupIndirectCommandNV-None-02944  
  The current bound graphics pipeline, as well as the pipelines it may reference, **must** have been created with `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`
- [](#VUID-VkBindShaderGroupIndirectCommandNV-index-02945)VUID-VkBindShaderGroupIndirectCommandNV-index-02945  
  The `index` **must** be within range of the accessible shader groups of the current bound graphics pipeline. See [vkCmdBindPipelineShaderGroupNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindPipelineShaderGroupNV.html) for further details

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindShaderGroupIndirectCommandNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0