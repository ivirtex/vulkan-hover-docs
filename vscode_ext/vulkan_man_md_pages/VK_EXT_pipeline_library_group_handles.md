# VK_EXT_pipeline_library_group_handles(3) Manual Page

## Name

VK_EXT_pipeline_library_group_handles - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

499

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html)  
and  
[VK_KHR_pipeline_library](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_pipeline_library.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Hans-Kristian Arntzen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_pipeline_library_group_handles%5D%20@HansKristian-Work%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_pipeline_library_group_handles%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>HansKristian-Work</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_EXT_pipeline_library_group_handles](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_EXT_pipeline_library_group_handles.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

When using pipeline libraries in ray tracing pipelines, a library might
get linked into different pipelines in an incremental way. An
application can have a strategy where a ray tracing pipeline is
comprised of N pipeline libraries and is later augmented by creating a
new pipeline with N + 1 libraries. Without this extension, all group
handles must be re-queried as the group handle is tied to the pipeline,
not the library. This is problematic for applications which aim to
decouple construction of record buffers and the linkage of ray tracing
pipelines.

To aid in this, this extension enables support for querying group
handles directly from pipeline libraries. Group handles obtained from a
library **must** remain bitwise identical in any `VkPipeline` that links
to the library.

With this feature, the extension also improves compatibility with DXR
1.1 AddToStateObject(), which guarantees that group handles returned
remain bitwise identical between parent and child pipelines. In
addition, querying group handles from COLLECTION objects is also
supported with that API.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineLibraryGroupHandlesFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_PIPELINE_LIBRARY_GROUP_HANDLES_EXTENSION_NAME`

- `VK_EXT_PIPELINE_LIBRARY_GROUP_HANDLES_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_LIBRARY_GROUP_HANDLES_FEATURES_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-01-25 (Hans-Kristian Arntzen)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_pipeline_library_group_handles"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
