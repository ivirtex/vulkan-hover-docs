# VkShaderInfoTypeAMD(3) Manual Page

## Name

VkShaderInfoTypeAMD - Enum specifying which type of shader information to query



## [](#_c_specification)C Specification

Possible values of [vkGetShaderInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderInfoAMD.html)::`infoType`, specifying the information being queried from a shader, are:

```c++
// Provided by VK_AMD_shader_info
typedef enum VkShaderInfoTypeAMD {
    VK_SHADER_INFO_TYPE_STATISTICS_AMD = 0,
    VK_SHADER_INFO_TYPE_BINARY_AMD = 1,
    VK_SHADER_INFO_TYPE_DISASSEMBLY_AMD = 2,
} VkShaderInfoTypeAMD;
```

## [](#_description)Description

- `VK_SHADER_INFO_TYPE_STATISTICS_AMD` specifies that device resources used by a shader will be queried.
- `VK_SHADER_INFO_TYPE_BINARY_AMD` specifies that implementation-specific information will be queried.
- `VK_SHADER_INFO_TYPE_DISASSEMBLY_AMD` specifies that human-readable disassembly of a shader.

## [](#_see_also)See Also

[VK\_AMD\_shader\_info](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_shader_info.html), [vkGetShaderInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetShaderInfoAMD.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkShaderInfoTypeAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0