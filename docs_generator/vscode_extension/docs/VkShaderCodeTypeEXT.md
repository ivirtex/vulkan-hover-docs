# VkShaderCodeTypeEXT(3) Manual Page

## Name

VkShaderCodeTypeEXT - Indicate a shader code type



## [](#_c_specification)C Specification

Shader objects **can** be created using different types of shader code. Possible values of [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateInfoEXT.html)::`codeType`, are:

```c++
// Provided by VK_EXT_shader_object
typedef enum VkShaderCodeTypeEXT {
    VK_SHADER_CODE_TYPE_BINARY_EXT = 0,
    VK_SHADER_CODE_TYPE_SPIRV_EXT = 1,
} VkShaderCodeTypeEXT;
```

## [](#_description)Description

- `VK_SHADER_CODE_TYPE_BINARY_EXT` specifies shader code in an opaque, implementation-defined binary format specific to the physical device.
- `VK_SHADER_CODE_TYPE_SPIRV_EXT` specifies shader code in SPIR-V format.

## [](#_see_also)See Also

[VK\_EXT\_shader\_object](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_shader_object.html), [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkShaderCodeTypeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0