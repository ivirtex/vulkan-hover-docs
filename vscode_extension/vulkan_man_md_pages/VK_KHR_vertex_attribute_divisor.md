# VK_KHR_vertex_attribute_divisor(3) Manual Page

## Name

VK_KHR_vertex_attribute_divisor - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

526

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Shahbaz Youssefi <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_vertex_attribute_divisor%5D%20@syoussefi%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_vertex_attribute_divisor%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>syoussefi</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_vertex_attribute_divisor](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_vertex_attribute_divisor.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-09-20

**IP Status**  
No known IP claims.

**Contributors**  
- Shahbaz Youssefi, Google

- Contributors to
  [`VK_EXT_vertex_attribute_divisor`](VK_EXT_vertex_attribute_divisor.html)

## <a href="#_description" class="anchor"></a>Description

This extension is based on the
[`VK_EXT_vertex_attribute_divisor`](VK_EXT_vertex_attribute_divisor.html)
extension. The only difference is the new property
`supportsNonZeroFirstInstance`, which indicates support for non-zero
values in `firstInstance`. This allows the extension to be supported on
implementations that have traditionally only supported OpenGL ES.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkVertexInputBindingDivisorDescriptionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDivisorDescriptionKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorFeaturesKHR.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR.html)

- Extending
  [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html):

  - [VkPipelineVertexInputDivisorStateCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputDivisorStateCreateInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_VERTEX_ATTRIBUTE_DIVISOR_EXTENSION_NAME`

- `VK_KHR_VERTEX_ATTRIBUTE_DIVISOR_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_ATTRIBUTE_DIVISOR_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_PIPELINE_VERTEX_INPUT_DIVISOR_STATE_CREATE_INFO_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-09-20 (Shahbaz Youssefi)

  - First Version, based on
    [`VK_EXT_vertex_attribute_divisor`](VK_EXT_vertex_attribute_divisor.html)

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_vertex_attribute_divisor"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
