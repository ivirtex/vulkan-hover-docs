# VK_NV_glsl_shader(3) Manual Page

## Name

VK_NV_glsl_shader - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

13

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Deprecated* without replacement

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_glsl_shader%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_glsl_shader%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-02-14

**IP Status**  
No known IP claims.

**Contributors**  
- Piers Daniell, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows GLSL shaders written to the `GL_KHR_vulkan_glsl`
extension specification to be used instead of SPIR-V. The implementation
will automatically detect whether the shader is SPIR-V or GLSL, and
compile it appropriately.

## <a href="#_deprecation" class="anchor"></a>Deprecation

Functionality in this extension is outside of the scope of Vulkan and is
better served by a compiler library such as
[glslang](https://github.com/KhronosGroup/glslang). No new
implementations will support this extension, so applications **should**
not use it.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_GLSL_SHADER_EXTENSION_NAME`

- `VK_NV_GLSL_SHADER_SPEC_VERSION`

- Extending [VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html):

  - `VK_ERROR_INVALID_SHADER_NV`

## <a href="#_examples" class="anchor"></a>Examples

**Example 1**

Passing in GLSL code

``` c
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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-02-14 (Piers Daniell)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_glsl_shader"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
