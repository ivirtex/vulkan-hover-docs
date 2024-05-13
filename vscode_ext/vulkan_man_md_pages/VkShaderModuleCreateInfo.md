# VkShaderModuleCreateInfo(3) Manual Page

## Name

VkShaderModuleCreateInfo - Structure specifying parameters of a newly
created shader module



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkShaderModuleCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkShaderModuleCreateInfo {
    VkStructureType              sType;
    const void*                  pNext;
    VkShaderModuleCreateFlags    flags;
    size_t                       codeSize;
    const uint32_t*              pCode;
} VkShaderModuleCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `codeSize` is the size, in bytes, of the code pointed to by `pCode`.

- `pCode` is a pointer to code that is used to create the shader module.
  The type and format of the code is determined from the content of the
  memory addressed by `pCode`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkShaderModuleCreateInfo-codeSize-08735"
  id="VUID-VkShaderModuleCreateInfo-codeSize-08735"></a>
  VUID-VkShaderModuleCreateInfo-codeSize-08735  
  If pCode is a pointer to SPIR-V code, `codeSize` **must** be a
  multiple of 4

- <a href="#VUID-VkShaderModuleCreateInfo-pCode-08736"
  id="VUID-VkShaderModuleCreateInfo-pCode-08736"></a>
  VUID-VkShaderModuleCreateInfo-pCode-08736  
  If pCode is a pointer to SPIR-V code, `pCode` **must** point to valid
  SPIR-V code, formatted and packed as described by the [Khronos SPIR-V
  Specification](#spirv-spec)

- <a href="#VUID-VkShaderModuleCreateInfo-pCode-08737"
  id="VUID-VkShaderModuleCreateInfo-pCode-08737"></a>
  VUID-VkShaderModuleCreateInfo-pCode-08737  
  If pCode is a pointer to SPIR-V code, `pCode` **must** adhere to the
  validation rules described by the [Validation Rules within a
  Module](#spirvenv-module-validation) section of the [SPIR-V
  Environment](#spirvenv-capabilities) appendix

- <a href="#VUID-VkShaderModuleCreateInfo-pCode-08738"
  id="VUID-VkShaderModuleCreateInfo-pCode-08738"></a>
  VUID-VkShaderModuleCreateInfo-pCode-08738  
  If pCode is a pointer to SPIR-V code, `pCode` **must** declare the
  `Shader` capability for SPIR-V code

- <a href="#VUID-VkShaderModuleCreateInfo-pCode-08739"
  id="VUID-VkShaderModuleCreateInfo-pCode-08739"></a>
  VUID-VkShaderModuleCreateInfo-pCode-08739  
  If pCode is a pointer to SPIR-V code, `pCode` **must** not declare any
  capability that is not supported by the API, as described by the
  [Capabilities](#spirvenv-module-validation) section of the [SPIR-V
  Environment](#spirvenv-capabilities) appendix

- <a href="#VUID-VkShaderModuleCreateInfo-pCode-08740"
  id="VUID-VkShaderModuleCreateInfo-pCode-08740"></a>
  VUID-VkShaderModuleCreateInfo-pCode-08740  
  If pCode is a pointer to SPIR-V code, and `pCode` declares any of the
  capabilities listed in the [SPIR-V
  Environment](#spirvenv-capabilities-table) appendix, one of the
  corresponding requirements **must** be satisfied

- <a href="#VUID-VkShaderModuleCreateInfo-pCode-08741"
  id="VUID-VkShaderModuleCreateInfo-pCode-08741"></a>
  VUID-VkShaderModuleCreateInfo-pCode-08741  
  If pCode is a pointer to SPIR-V code, `pCode` **must** not declare any
  SPIR-V extension that is not supported by the API, as described by the
  [Extension](#spirvenv-extensions) section of the [SPIR-V
  Environment](#spirvenv-capabilities) appendix

- <a href="#VUID-VkShaderModuleCreateInfo-pCode-08742"
  id="VUID-VkShaderModuleCreateInfo-pCode-08742"></a>
  VUID-VkShaderModuleCreateInfo-pCode-08742  
  If pCode is a pointer to SPIR-V code, and `pCode` declares any of the
  SPIR-V extensions listed in the [SPIR-V
  Environment](#spirvenv-extensions-table) appendix, one of the
  corresponding requirements **must** be satisfied

- <a href="#VUID-VkShaderModuleCreateInfo-pCode-07912"
  id="VUID-VkShaderModuleCreateInfo-pCode-07912"></a>
  VUID-VkShaderModuleCreateInfo-pCode-07912  
  If the [VK_NV_glsl_shader](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_glsl_shader.html) extension is not
  enabled, `pCode` **must** be a pointer to SPIR-V code

- <a href="#VUID-VkShaderModuleCreateInfo-pCode-01379"
  id="VUID-VkShaderModuleCreateInfo-pCode-01379"></a>
  VUID-VkShaderModuleCreateInfo-pCode-01379  
  If `pCode` is a pointer to GLSL code, it **must** be valid GLSL code
  written to the `GL_KHR_vulkan_glsl` GLSL extension specification

- <a href="#VUID-VkShaderModuleCreateInfo-codeSize-01085"
  id="VUID-VkShaderModuleCreateInfo-codeSize-01085"></a>
  VUID-VkShaderModuleCreateInfo-codeSize-01085  
  `codeSize` **must** be greater than 0

Valid Usage (Implicit)

- <a href="#VUID-VkShaderModuleCreateInfo-sType-sType"
  id="VUID-VkShaderModuleCreateInfo-sType-sType"></a>
  VUID-VkShaderModuleCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO`

- <a href="#VUID-VkShaderModuleCreateInfo-flags-zerobitmask"
  id="VUID-VkShaderModuleCreateInfo-flags-zerobitmask"></a>
  VUID-VkShaderModuleCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-VkShaderModuleCreateInfo-pCode-parameter"
  id="VUID-VkShaderModuleCreateInfo-pCode-parameter"></a>
  VUID-VkShaderModuleCreateInfo-pCode-parameter  
  `pCode` **must** be a valid pointer to an array of 4codeSize​
  `uint32_t` values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkShaderModuleCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateShaderModule.html),
[vkGetShaderModuleCreateInfoIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderModuleCreateInfoIdentifierEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShaderModuleCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
