# VK\_NV\_glsl\_shader(3) Manual Page

## Name

VK\_NV\_glsl\_shader - device extension



## [](#_registered_extension_number)Registered Extension Number

13

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Deprecated* without replacement

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_glsl_shader%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_glsl_shader%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-02-14

**IP Status**

No known IP claims.

**Contributors**

- Piers Daniell, NVIDIA

## [](#_description)Description

This extension allows GLSL shaders written to the `GL_KHR_vulkan_glsl` extension specification to be used instead of SPIR-V. The implementation will automatically detect whether the shader is SPIR-V or GLSL, and compile it appropriately.

## [](#_deprecation)Deprecation

Functionality in this extension is outside of the scope of Vulkan and is better served by a compiler library such as [glslang](https://github.com/KhronosGroup/glslang). No new implementations will support this extension, so applications **should** not use it.

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_GLSL_SHADER_EXTENSION_NAME`
- `VK_NV_GLSL_SHADER_SPEC_VERSION`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_INVALID_SHADER_NV`

## [](#_examples)Examples

**Example 1**

Passing in GLSL code

```c++
    char const vss[] =
        "#version 450 core\n"
        "layout(location = 0) in vec2 aVertex;\n"
        "layout(location = 1) in vec4 aColor;\n"
        "out vec4 vColor;\n"
        "void main()\n"
        "{\n"
        "    vColor = aColor;\n"
        "    gl_Position = vec4(aVertex, 0, 1);\n"
        "}\n"
    ;
    VkShaderModuleCreateInfo vertexShaderInfo = { VK_STRUCTURE_TYPE_SHADER_MODULE_CREATE_INFO };
    vertexShaderInfo.codeSize = sizeof vss;
    vertexShaderInfo.pCode = vss;
    VkShaderModule vertexShader;
    vkCreateShaderModule(device, &vertexShaderInfo, 0, &vertexShader);
```

## [](#_version_history)Version History

- Revision 1, 2016-02-14 (Piers Daniell)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_glsl_shader).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0