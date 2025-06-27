# VkShaderInfoTypeAMD(3) Manual Page

## Name

VkShaderInfoTypeAMD - Enum specifying which type of shader information
to query



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of
[vkGetShaderInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderInfoAMD.html)::`infoType`, specifying
the information being queried from a shader, are:

``` c
// Provided by VK_AMD_shader_info
typedef enum VkShaderInfoTypeAMD {
    VK_SHADER_INFO_TYPE_STATISTICS_AMD = 0,
    VK_SHADER_INFO_TYPE_BINARY_AMD = 1,
    VK_SHADER_INFO_TYPE_DISASSEMBLY_AMD = 2,
} VkShaderInfoTypeAMD;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SHADER_INFO_TYPE_STATISTICS_AMD` specifies that device resources
  used by a shader will be queried.

- `VK_SHADER_INFO_TYPE_BINARY_AMD` specifies that
  implementation-specific information will be queried.

- `VK_SHADER_INFO_TYPE_DISASSEMBLY_AMD` specifies that human-readable
  disassembly of a shader.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_shader_info](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_shader_info.html),
[vkGetShaderInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderInfoAMD.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkShaderInfoTypeAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
