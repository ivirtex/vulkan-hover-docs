# VK\_AMD\_shader\_fragment\_mask(3) Manual Page

## Name

VK\_AMD\_shader\_fragment\_mask - device extension



## [](#_registered_extension_number)Registered Extension Number

138

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_AMD\_shader\_fragment\_mask](https://github.khronos.org/SPIRV-Registry/extensions/AMD/SPV_AMD_shader_fragment_mask.html)

## [](#_contact)Contact

- Aaron Hagan [\[GitHub\]AaronHaganAMD](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_shader_fragment_mask%5D%20%40AaronHaganAMD%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_shader_fragment_mask%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-08-16

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_AMD_shader_fragment_mask`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/amd/GL_AMD_shader_fragment_mask.txt)

**Contributors**

- Aaron Hagan, AMD
- Daniel Rakos, AMD
- Timothy Lottes, AMD

## [](#_description)Description

This extension provides efficient read access to the fragment mask in compressed multisampled color surfaces. The fragment mask is a lookup table that associates color samples with color fragment values.

From a shader, the fragment mask can be fetched with a call to `fragmentMaskFetchAMD`, which returns a single `uint` where each subsequent four bits specify the color fragment index corresponding to the color sample, starting from the least significant bit. For example, when eight color samples are used, the color fragment index for color sample 0 will be in bits 0-3 of the fragment mask, for color sample 7 the index will be in bits 28-31.

The color fragment for a particular color sample may then be fetched with the corresponding fragment mask value using the `fragmentFetchAMD` shader function.

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_SHADER_FRAGMENT_MASK_EXTENSION_NAME`
- `VK_AMD_SHADER_FRAGMENT_MASK_SPEC_VERSION`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`FragmentMaskAMD`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-FragmentMaskAMD)

## [](#_examples)Examples

This example shows a shader that queries the fragment mask from a multisampled compressed surface and uses it to query fragment values.

```c++
#version 450 core

#extension GL_AMD_shader_fragment_mask: enable

layout(binding = 0) uniform sampler2DMS       s2DMS;
layout(binding = 1) uniform isampler2DMSArray is2DMSArray;

layout(binding = 2, input_attachment_index = 0) uniform usubpassInputMS usubpassMS;

layout(location = 0) out vec4 fragColor;

void main()
{
    vec4 fragOne = vec4(0.0);

    uint fragMask = fragmentMaskFetchAMD(s2DMS, ivec2(2, 3));
    uint fragIndex = (fragMask & 0xF0) >> 4;
    fragOne += fragmentFetchAMD(s2DMS, ivec2(2, 3), 1);

    fragMask = fragmentMaskFetchAMD(is2DMSArray, ivec3(2, 3, 1));
    fragIndex = (fragMask & 0xF0) >> 4;
    fragOne += fragmentFetchAMD(is2DMSArray, ivec3(2, 3, 1), fragIndex);

    fragMask = fragmentMaskFetchAMD(usubpassMS);
    fragIndex = (fragMask & 0xF0) >> 4;
    fragOne += fragmentFetchAMD(usubpassMS, fragIndex);

    fragColor = fragOne;
}
```

## [](#_version_history)Version History

- Revision 1, 2017-08-16 (Aaron Hagan)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_shader_fragment_mask)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0