# VK_KHR_pipeline_executable_properties(3) Manual Page

## Name

VK_KHR_pipeline_executable_properties - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

270

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">Developer tools</a>

## <a href="#_contact" class="anchor"></a>Contact

- Faith Ekstrand <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_pipeline_executable_properties%5D%20@gfxstrand%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_pipeline_executable_properties%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>gfxstrand</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-05-28

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
**Contributors**  
- Faith Ekstrand, Intel

- Ian Romanick, Intel

- Kenneth Graunke, Intel

- Baldur Karlsson, Valve

- Jesse Hall, Google

- Jeff Bolz, Nvidia

- Piers Daniel, Nvidia

- Tobias Hector, AMD

- Jan-Harald Fredriksen, ARM

- Tom Olson, ARM

- Daniel Koch, Nvidia

- Spencer Fricke, Samsung

## <a href="#_description" class="anchor"></a>Description

When a pipeline is created, its state and shaders are compiled into zero
or more device-specific executables, which are used when executing
commands against that pipeline. This extension adds a mechanism to query
properties and statistics about the different executables produced by
the pipeline compilation process. This is intended to be used by
debugging and performance tools to allow them to provide more detailed
information to the user. Certain compile time shader statistics provided
through this extension may be useful to developers for debugging or
performance analysis.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetPipelineExecutableInternalRepresentationsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutableInternalRepresentationsKHR.html)

- [vkGetPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutablePropertiesKHR.html)

- [vkGetPipelineExecutableStatisticsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineExecutableStatisticsKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkPipelineExecutableInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableInfoKHR.html)

- [VkPipelineExecutableInternalRepresentationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableInternalRepresentationKHR.html)

- [VkPipelineExecutablePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutablePropertiesKHR.html)

- [VkPipelineExecutableStatisticKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableStatisticKHR.html)

- [VkPipelineInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInfoKHR.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineExecutablePropertiesFeaturesKHR.html)

## <a href="#_new_unions" class="anchor"></a>New Unions

- [VkPipelineExecutableStatisticValueKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableStatisticValueKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkPipelineExecutableStatisticFormatKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineExecutableStatisticFormatKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_PIPELINE_EXECUTABLE_PROPERTIES_EXTENSION_NAME`

- `VK_KHR_PIPELINE_EXECUTABLE_PROPERTIES_SPEC_VERSION`

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_CAPTURE_INTERNAL_REPRESENTATIONS_BIT_KHR`

  - `VK_PIPELINE_CREATE_CAPTURE_STATISTICS_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_EXECUTABLE_PROPERTIES_FEATURES_KHR`

  - `VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_INTERNAL_REPRESENTATION_KHR`

  - `VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_PIPELINE_EXECUTABLE_STATISTIC_KHR`

  - `VK_STRUCTURE_TYPE_PIPELINE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) What should we call the pieces of the pipeline which are produced by
the compilation process and about which you can query properties and
statistics?

**RESOLVED**: Call them “executables”. The name “binary” was used in
early drafts of the extension but it was determined that “pipeline
binary” could have a fairly broad meaning (such as a binary serialized
form of an entire pipeline) and was too big of a namespace for the very
specific needs of this extension.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-05-28 (Faith Ekstrand)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_pipeline_executable_properties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
