# VK_AMD_shader_fragment_mask(3) Manual Page

## Name

VK_AMD_shader_fragment_mask - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

138

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_AMD_shader_fragment_mask](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/AMD/SPV_AMD_shader_fragment_mask.html)

## <a href="#_contact" class="anchor"></a>Contact

- Aaron Hagan <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_shader_fragment_mask%5D%20@AaronHaganAMD%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_shader_fragment_mask%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>AaronHaganAMD</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-08-16

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_AMD_shader_fragment_mask`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/amd/GL_AMD_shader_fragment_mask.txt)

**Contributors**  
- Aaron Hagan, AMD

- Daniel Rakos, AMD

- Timothy Lottes, AMD

## <a href="#_description" class="anchor"></a>Description

This extension provides efficient read access to the fragment mask in
compressed multisampled color surfaces. The fragment mask is a lookup
table that associates color samples with color fragment values.

From a shader, the fragment mask can be fetched with a call to
`fragmentMaskFetchAMD`, which returns a single `uint` where each
subsequent four bits specify the color fragment index corresponding to
the color sample, starting from the least significant bit. For example,
when eight color samples are used, the color fragment index for color
sample 0 will be in bits 0-3 of the fragment mask, for color sample 7
the index will be in bits 28-31.

The color fragment for a particular color sample may then be fetched
with the corresponding fragment mask value using the `fragmentFetchAMD`
shader function.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_SHADER_FRAGMENT_MASK_EXTENSION_NAME`

- `VK_AMD_SHADER_FRAGMENT_MASK_SPEC_VERSION`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-FragmentMaskAMD"
  target="_blank" rel="noopener"><code>FragmentMaskAMD</code></a>

## <a href="#_examples" class="anchor"></a>Examples

This example shows a shader that queries the fragment mask from a
multisampled compressed surface and uses it to query fragment values.

``` c
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

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-08-16 (Aaron Hagan)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_shader_fragment_mask"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
