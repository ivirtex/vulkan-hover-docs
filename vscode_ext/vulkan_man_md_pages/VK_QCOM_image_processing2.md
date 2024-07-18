# VK_QCOM_image_processing2(3) Manual Page

## Name

VK_QCOM_image_processing2 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

519

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_QCOM_image_processing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_image_processing.html)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_QCOM_image_processing2](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/QCOM/SPV_QCOM_image_processing2.html)

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_image_processing2%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_image_processing2%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-03-10

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_QCOM_image_processing2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/qcom/GLSL_QCOM_image_processing2.txt)

**Contributors**  
- Jeff Leger, Qualcomm Technologies, Inc.

## <a href="#_description" class="anchor"></a>Description

This extension enables support for the SPIR-V `TextureBlockMatch2QCOM`
capability. It builds on the functionality of QCOM_image_processing with
the addition of 4 new image processing operations.

- The `opImageBlockMatchWindowSADQCOM`\` SPIR-V instruction builds upon
  the functionality of `opImageBlockMatchSADQCOM`\` by repeatedly
  performing block match operations across a 2D window. The “2D
  windowExtent” and “compareMode” are specified by
  [VkSamplerBlockMatchWindowCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerBlockMatchWindowCreateInfoQCOM.html)
  in the sampler used to create the *target image*. Like
  `OpImageBlockMatchSADQCOM`, `opImageBlockMatchWindowSADQCOM` computes
  an error metric, that describes whether a block of texels in the
  *target image* matches a corresponding block of texels in the
  *reference image*. Unlike `OpImageBlockMatchSADQCOM`, this instruction
  computes an error metric at each (X,Y) location within the 2D window
  and returns either the minimum or maximum error. The instruction only
  supports single-component formats. Refer to the pseudocode below for
  details.

- The `opImageBlockMatchWindowSSDQCOM` follows the same pattern,
  computing the SSD error metric at each location within the 2D window.

- The `opImageBlockMatchGatherSADQCOM` builds upon
  `OpImageBlockMatchSADQCOM`. This instruction computes an error metric,
  that describes whether a block of texels in the *target image* matches
  a corresponding block of texels in the *reference image*. The
  instruction computes the SAD error metric at 4 texel offsets and
  returns the error metric for each offset in the X,Y,Z,and W
  components. The instruction only supports single-component texture
  formats. Refer to the pseudocode below for details.

- The `opImageBlockMatchGatherSSDQCOM` follows the same pattern,
  computing the SSD error metric for 4 offsets.

Each of the above 4 image processing instructions are limited to
single-component formats.

Below is the pseudocode for GLSL built-in function
`textureWindowBlockMatchSADQCOM`. The pseudocode for
`textureWindowBlockMatchSSD` is identical other than replacing all
instances of `"SAD"` with `"SSD"`.

``` c
vec4 textureBlockMatchWindowSAD( sampler2D target,
                                 uvec2 targetCoord,
                                 samler2D reference,
                                 uvec2 refCoord,
                                 uvec2 blocksize) {
    // compareMode (MIN or MAX) comes from the vkSampler associated with `target`
    // uvec2 window  comes from the vkSampler associated with `target`
    minSAD = INF;
    maxSAD = -INF;
    uvec2 minCoord;
    uvec2 maxCoord;

    for (uint x=0, x < window.width; x++) {
        for (uint y=0; y < window.height; y++) {
            float SAD = textureBlockMatchSAD(target,
                                            targetCoord + uvec2(x, y),
                                            reference,
                                            refCoord,
                                            blocksize).x;
            // Note: the below comparison operator will produce undefined results
            // if SAD is a denorm value.
            if (SAD < minSAD) {
                minSAD = SAD;
                minCoord = uvec2(x,y);
            }
            if (SAD > maxSAD) {
                maxSAD = SAD;
                maxCoord = uvec2(x,y);
            }
        }
    }
    if (compareMode=MIN) {
        return vec4(minSAD, minCoord.x, minCoord.y, 0.0);
    } else {
        return vec4(maxSAD, maxCoord.x, maxCoord.y, 0.0);
    }
}
```

Below is the pseudocode for `textureBlockMatchGatherSADQCOM`. The
pseudocode for `textureBlockMatchGatherSSD` follows an identical
pattern.

``` c
vec4 textureBlockMatchGatherSAD( sampler2D target,
                                 uvec2 targetCoord,
                                 samler2D reference,
                                 uvec2 refCoord,
                                 uvec2 blocksize) {
    vec4 out;
    for (uint x=0, x<4; x++) {
            float SAD = textureBlockMatchSAD(target,
                                            targetCoord + uvec2(x, 0),
                                            reference,
                                            refCoord,
                                            blocksize).x;
            out[x] = SAD;
    }
    return out;
}
```

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceImageProcessing2FeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageProcessing2FeaturesQCOM.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceImageProcessing2PropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageProcessing2PropertiesQCOM.html)

- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html):

  - [VkSamplerBlockMatchWindowCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerBlockMatchWindowCreateInfoQCOM.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlockMatchWindowCompareModeQCOM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_IMAGE_PROCESSING_2_EXTENSION_NAME`

- `VK_QCOM_IMAGE_PROCESSING_2_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_2_FEATURES_QCOM`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_2_PROPERTIES_QCOM`

  - `VK_STRUCTURE_TYPE_SAMPLER_BLOCK_MATCH_WINDOW_CREATE_INFO_QCOM`

## <a href="#_issues" class="anchor"></a>Issues

1\) What is the precision of the min/max comparison checks?

**RESOLVED**: Intermediate computations for the new operations are
performed at 16-bit floating-point precision. If the value of
`"float SAD"` in the above code sample is a 16-bit denorm value, then
behavior of the MIN/MAX comparison is undefined.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-03-10 (Jeff Leger)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_image_processing2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
