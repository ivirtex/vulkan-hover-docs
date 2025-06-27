# VK_EXT_pipeline_properties(3) Manual Page

## Name

VK_EXT_pipeline_properties - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

373

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Mukund Keshava <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_pipeline_properties%5D%20@mkeshavanv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_pipeline_properties%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mkeshavanv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-04-19

**IP Status**  
No known IP claims.

**Contributors**  
- Mukund Keshava, NVIDIA

- Daniel Koch, NVIDIA

- Mark Bellamy, Arm

## <a href="#_description" class="anchor"></a>Description

Vulkan SC requires offline compilation of pipelines. In order to support
this, the pipeline state is represented in a [JSON
schema](https://github.com/KhronosGroup/VulkanSC-Docs/wiki/JSON-schema)
that is read by an offline tool for compilation.

One method of developing a Vulkan SC application is to author a Vulkan
application and use a layer to record and serialize the pipeline state
and shaders for offline compilation. Each pipeline is represented by a
separate JSON file, and can be identified with a `pipelineIdentifier`.

Once the pipelines have been compiled by the offline pipeline cache
compiler, the Vulkan SC application can then use this
`pipelineIdentifier` for identifying the pipeline via Vulkan SC’s
`VkPipelineIdentifierInfo` structure.

This extension allows the Vulkan application to query the
`pipelineIdentifier` associated with each pipeline so that the
application can store this with its pipeline metadata and the Vulkan SC
application will then use to map the same state to an entry in the
Vulkan SC pipeline cache.

It is expected that this extension will initially be implemented in the
json generation layer, although we can envision that there might be
future uses for it in native Vulkan drivers as well.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetPipelinePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelinePropertiesEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkPipelineInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInfoEXT.html)

- [VkPipelinePropertiesIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelinePropertiesIdentifierEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePipelinePropertiesFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelinePropertiesFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_PIPELINE_PROPERTIES_EXTENSION_NAME`

- `VK_EXT_PIPELINE_PROPERTIES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_PROPERTIES_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_INFO_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_PROPERTIES_IDENTIFIER_EXT`

## <a href="#_issues" class="anchor"></a>Issues

\(1\) This extension does not make sense on a strict Vulkan SC
implementation. It may however be of potential use in a non-strict
Vulkan SC implementation. Should this extension be enabled as part of
Vulkan SC as well?

**RESOLVED**: No. This extension will not be enabled for Vulkan SC.

\(2\) This is intended to be a general pipeline properties query, but is
currently only retrieving the pipeline identifier. Should the pipeline
identifier query be mandatory for this extension and for all queries
using this command?

**RESOLVED**: Use [VkBaseOutStructure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBaseOutStructure.html) for the
return parameter. Currently this is required to actually be a
[VkPipelinePropertiesIdentifierEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelinePropertiesIdentifierEXT.html)
structure, but that could be relaxed in the future to allow other
structure types or to allow other structures to be chained in along with
this one.

\(3\) Should there be a feature structure? Should it be required?

**RESOLVED**: Add a feature structure, and a feature for querying
pipeline identifier, but allow it to be optional so that this extension
can be used as the basis for other pipeline property queries without
requiring the pipeline identifier to be supported.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-04-19 (Mukund Keshava, Daniel Koch)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_pipeline_properties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
