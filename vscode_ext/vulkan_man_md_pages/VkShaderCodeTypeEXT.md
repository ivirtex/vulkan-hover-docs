# VkShaderCodeTypeEXT(3) Manual Page

## Name

VkShaderCodeTypeEXT - Indicate a shader code type



## <a href="#_c_specification" class="anchor"></a>C Specification

Shader objects **can** be created using different types of shader code.
Possible values of
[VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html)::`codeType`, are:

``` c
// Provided by VK_EXT_shader_object
typedef enum VkShaderCodeTypeEXT {
    VK_SHADER_CODE_TYPE_BINARY_EXT = 0,
    VK_SHADER_CODE_TYPE_SPIRV_EXT = 1,
} VkShaderCodeTypeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SHADER_CODE_TYPE_BINARY_EXT` specifies shader code in an opaque,
  implementation-defined binary format specific to the physical device.

- `VK_SHADER_CODE_TYPE_SPIRV_EXT` specifies shader code in SPIR-V
  format.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_object](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_object.html),
[VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShaderCodeTypeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
