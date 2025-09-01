# VK\_QCOM\_image\_processing2(3) Manual Page

## Name

VK\_QCOM\_image\_processing2 - device extension



## [](#_registered_extension_number)Registered Extension Number

519

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_QCOM\_image\_processing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_image_processing.html)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_QCOM\_image\_processing2](https://github.khronos.org/SPIRV-Registry/extensions/QCOM/SPV_QCOM_image_processing2.html)

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_image_processing2%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_image_processing2%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-03-10

**Interactions and External Dependencies**

- This extension provides API support for [`GL_QCOM_image_processing2`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/qcom/GLSL_QCOM_image_processing2.txt)

**Contributors**

- Jeff Leger, Qualcomm Technologies, Inc.

## [](#_description)Description

This extension enables support for the SPIR-V `TextureBlockMatch2QCOM` capability. It builds on the functionality of QCOM\_image\_processing with the addition of 4 new image processing operations.

- The `opImageBlockMatchWindowSADQCOM`\` SPIR-V instruction builds upon the functionality of `opImageBlockMatchSADQCOM`\` by repeatedly performing block match operations across a 2D window. The “2D windowExtent” and “compareMode” are specified by [VkSamplerBlockMatchWindowCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerBlockMatchWindowCreateInfoQCOM.html) in the sampler used to create the *target image*. Like `OpImageBlockMatchSADQCOM`, `opImageBlockMatchWindowSADQCOM` computes an error metric, that describes whether a block of texels in the *target image* matches a corresponding block of texels in the *reference image*. Unlike `OpImageBlockMatchSADQCOM`, this instruction computes an error metric at each (X,Y) location within the 2D window and returns either the minimum or maximum error. The instruction only supports single-component formats. Refer to the pseudocode below for details.
- The `opImageBlockMatchWindowSSDQCOM` follows the same pattern, computing the SSD error metric at each location within the 2D window.
- The `opImageBlockMatchGatherSADQCOM` builds upon `OpImageBlockMatchSADQCOM`. This instruction computes an error metric, that describes whether a block of texels in the *target image* matches a corresponding block of texels in the *reference image*. The instruction computes the SAD error metric at 4 texel offsets and returns the error metric for each offset in the X,Y,Z,and W components. The instruction only supports single-component texture formats. Refer to the pseudocode below for details.
- The `opImageBlockMatchGatherSSDQCOM` follows the same pattern, computing the SSD error metric for 4 offsets.

Each of the above 4 image processing instructions are limited to single-component formats.

Below is the pseudocode for GLSL built-in function `textureWindowBlockMatchSADQCOM`. The pseudocode for `textureWindowBlockMatchSSD` is identical other than replacing all instances of `"SAD"` with `"SSD"`.

```c
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

Below is the pseudocode for `textureBlockMatchGatherSADQCOM`. The pseudocode for `textureBlockMatchGatherSSD` follows an identical pattern.

```c
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

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceImageProcessing2FeaturesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageProcessing2FeaturesQCOM.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceImageProcessing2PropertiesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageProcessing2PropertiesQCOM.html)
- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html):
  
  - [VkSamplerBlockMatchWindowCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerBlockMatchWindowCreateInfoQCOM.html)

## [](#_new_enums)New Enums

- [VkBlockMatchWindowCompareModeQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlockMatchWindowCompareModeQCOM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_IMAGE_PROCESSING_2_EXTENSION_NAME`
- `VK_QCOM_IMAGE_PROCESSING_2_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_2_FEATURES_QCOM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGE_PROCESSING_2_PROPERTIES_QCOM`
  - `VK_STRUCTURE_TYPE_SAMPLER_BLOCK_MATCH_WINDOW_CREATE_INFO_QCOM`

## [](#_issues)Issues

1\) What is the precision of the min/max comparison checks?

**RESOLVED**: Intermediate computations for the new operations are performed at 16-bit floating-point precision. If the value of `"float SAD"` in the above code sample is a 16-bit denorm value, then behavior of the MIN/MAX comparison is undefined.

## [](#_version_history)Version History

- Revision 1, 2023-03-10 (Jeff Leger)

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_image_processing2).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0