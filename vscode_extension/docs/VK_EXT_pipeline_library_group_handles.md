# VK\_EXT\_pipeline\_library\_group\_handles(3) Manual Page

## Name

VK\_EXT\_pipeline\_library\_group\_handles - device extension



## [](#_registered_extension_number)Registered Extension Number

499

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html)  
and  
[VK\_KHR\_pipeline\_library](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_library.html)

## [](#_contact)Contact

- Hans-Kristian Arntzen [\[GitHub\]HansKristian-Work](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_pipeline_library_group_handles%5D%20%40HansKristian-Work%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_pipeline_library_group_handles%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_EXT\_pipeline\_library\_group\_handles](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_pipeline_library_group_handles.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-01-25

**IP Status**

No known IP claims.

**Contributors**

- Hans-Kristian Arntzen, Valve
- Stuart Smith, AMD
- Ricardo Garcia, Igalia
- Lionel Landwerlin, Intel
- Eric Werness, NVIDIA
- Daniel Koch, NVIDIA

## [](#_description)Description

When using pipeline libraries in ray tracing pipelines, a library might get linked into different pipelines in an incremental way. An application can have a strategy where a ray tracing pipeline is comprised of N pipeline libraries and is later augmented by creating a new pipeline with N + 1 libraries. Without this extension, all group handles must be re-queried as the group handle is tied to the pipeline, not the library. This is problematic for applications which aim to decouple construction of record buffers and the linkage of ray tracing pipelines.

To aid in this, this extension enables support for querying group handles directly from pipeline libraries. Group handles obtained from a library **must** remain bitwise identical in any `VkPipeline` that links to the library.

With this feature, the extension also improves compatibility with DXR 1.1 AddToStateObject(), which guarantees that group handles returned remain bitwise identical between parent and child pipelines. In addition, querying group handles from COLLECTION objects is also supported with that API.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_PIPELINE_LIBRARY_GROUP_HANDLES_EXTENSION_NAME`
- `VK_EXT_PIPELINE_LIBRARY_GROUP_HANDLES_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_LIBRARY_GROUP_HANDLES_FEATURES_EXT`

## [](#_version_history)Version History

- Revision 1, 2023-01-25 (Hans-Kristian Arntzen)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_pipeline_library_group_handles)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0