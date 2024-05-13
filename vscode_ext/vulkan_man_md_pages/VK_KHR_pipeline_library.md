# VK_KHR_pipeline_library(3) Manual Page

## Name

VK_KHR_pipeline_library - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

291

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Christoph Kubisch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_pipeline_library%5D%20@pixeljetstream%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_pipeline_library%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pixeljetstream</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-01-08

**IP Status**  
No known IP claims.

**Contributors**  
- See contributors to
  [`VK_KHR_ray_tracing_pipeline`](VK_KHR_ray_tracing_pipeline.html)

## <a href="#_description" class="anchor"></a>Description

A pipeline library is a special pipeline that cannot be bound, instead
it defines a set of shaders and shader groups which can be linked into
other pipelines. This extension defines the infrastructure for pipeline
libraries, but does not specify the creation or usage of pipeline
libraries. This is left to additional dependent extensions.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html):

  - [VkPipelineLibraryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLibraryCreateInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_PIPELINE_LIBRARY_EXTENSION_NAME`

- `VK_KHR_PIPELINE_LIBRARY_SPEC_VERSION`

- Extending [VkPipelineCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlagBits.html):

  - `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PIPELINE_LIBRARY_CREATE_INFO_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-01-08 (Christoph Kubisch)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_pipeline_library"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
