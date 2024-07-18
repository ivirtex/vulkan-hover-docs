# VkBindShaderGroupIndirectCommandNV(3) Manual Page

## Name

VkBindShaderGroupIndirectCommandNV - Structure specifying input data for
a single shader group command token



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindShaderGroupIndirectCommandNV` structure specifies the input
data for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_SHADER_GROUP_NV` token.

``` c
// Provided by VK_NV_device_generated_commands
typedef struct VkBindShaderGroupIndirectCommandNV {
    uint32_t    groupIndex;
} VkBindShaderGroupIndirectCommandNV;
```

## <a href="#_members" class="anchor"></a>Members

- `groupIndex` specifies which shader group of the current bound
  graphics pipeline is used.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBindShaderGroupIndirectCommandNV-None-02944"
  id="VUID-VkBindShaderGroupIndirectCommandNV-None-02944"></a>
  VUID-VkBindShaderGroupIndirectCommandNV-None-02944  
  The current bound graphics pipeline, as well as the pipelines it may
  reference, **must** have been created with
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`

- <a href="#VUID-VkBindShaderGroupIndirectCommandNV-index-02945"
  id="VUID-VkBindShaderGroupIndirectCommandNV-index-02945"></a>
  VUID-VkBindShaderGroupIndirectCommandNV-index-02945  
  The `index` **must** be within range of the accessible shader groups
  of the current bound graphics pipeline. See
  [vkCmdBindPipelineShaderGroupNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindPipelineShaderGroupNV.html)
  for further details

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindShaderGroupIndirectCommandNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
