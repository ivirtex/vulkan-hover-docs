# VK_QCOM_filter_cubic_clamp(3) Manual Page

## Name

VK_QCOM_filter_cubic_clamp - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

522

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_filter_cubic](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_filter_cubic.html)  
and  
     [Version 1.2](#versions-1.2)  
     or  
     [VK_EXT_sampler_filter_minmax](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sampler_filter_minmax.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_filter_cubic_clamp%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_filter_cubic_clamp%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-08-02

**Contributors**  
- Jeff Leger, Qualcomm Technologies, Inc.

## <a href="#_description" class="anchor"></a>Description

This extension extends cubic filtering by adding the ability to enable
an anti-ringing clamp. Cubic filtering samples from a 4x4 region of
texels and computes a cubic weighted average of the region. In some
cases, the resulting value is outside the range of any of the texels in
the 4x4 region. This is sometimes referred to as “filter overshoot” or
“filter ringing” and can occur when there is a sharp discontinuity in
the 4x4 region being filtered. For some use cases this “ringing” can
produces unacceptable artifacts.

The solution to the ringing problem is to clamp the post-cubic-filtered
value to be within the max and min of texel values in the 4x4 region.
While such “range clamping” can be performed in shader code, the
additional texture fetches and clamping ALU operations can be costly.

Certain Adreno GPUs are able to perform the range clamp in the texture
unit during cubic filtering at significant performance/power savings
versus a shader-based clamping approach. This extension exposes such
hardware functionality.

This extension extends
[VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html), adding
`VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM` which
enables the range clamp operation.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceCubicClampFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCubicClampFeaturesQCOM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_FILTER_CUBIC_CLAMP_EXTENSION_NAME`

- `VK_QCOM_FILTER_CUBIC_CLAMP_SPEC_VERSION`

- Extending [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html):

  - `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUBIC_CLAMP_FEATURES_QCOM`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-08-02 (jleger)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_filter_cubic_clamp"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
