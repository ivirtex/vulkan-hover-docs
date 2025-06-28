# VK\_EXT\_provoking\_vertex(3) Manual Page

## Name

VK\_EXT\_provoking\_vertex - device extension



## [](#_registered_extension_number)Registered Extension Number

255

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_special_use)Special Use

- [OpenGL / ES support](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Jesse Hall [\[GitHub\]jessehall](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_provoking_vertex%5D%20%40jessehall%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_provoking_vertex%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-02-22

**IP Status**

No known IP claims.

**Contributors**

- Alexis Hétu, Google
- Bill Licea-Kane, Qualcomm
- Daniel Koch, Nvidia
- Jamie Madill, Google
- Jan-Harald Fredriksen, Arm
- Faith Ekstrand, Intel
- Jeff Bolz, Nvidia
- Jeff Leger, Qualcomm
- Jesse Hall, Google
- Jörg Wagner, Arm
- Matthew Netsch, Qualcomm
- Mike Blumenkrantz, Valve
- Piers Daniell, Nvidia
- Tobias Hector, AMD

## [](#_description)Description

This extension allows changing the provoking vertex convention between Vulkan’s default convention (first vertex) and OpenGL’s convention (last vertex).

This extension is intended for use by API-translation layers that implement APIs like OpenGL on top of Vulkan, and need to match the source API’s provoking vertex convention. Applications using Vulkan directly should use Vulkan’s default convention.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceProvokingVertexFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProvokingVertexFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceProvokingVertexPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProvokingVertexPropertiesEXT.html)
- Extending [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html):
  
  - [VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkProvokingVertexModeEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_PROVOKING_VERTEX_EXTENSION_NAME`
- `VK_EXT_PROVOKING_VERTEX_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROVOKING_VERTEX_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROVOKING_VERTEX_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_PROVOKING_VERTEX_STATE_CREATE_INFO_EXT`

## [](#_issues)Issues

1\) At what granularity should this state be set?

**RESOLVED**: At pipeline bind, with an optional per-render pass restriction.

The most natural place to put this state is in the graphics pipeline object. Some implementations require it to be known when creating the pipeline, and pipeline state is convenient for implementing OpenGL 3.2’s glProvokingVertex, which can change the state between draw calls. However, some implementations can only change it approximately render pass granularity. To accommodate both, provoking vertex will be pipeline state, but implementations can require that only one mode is used within a render pass instance; the render pass’s mode is chosen implicitly when the first pipeline is bound.

2\) Does the provoking vertex mode affect the order that vertices are written to transform feedback buffers?

**RESOLVED**: Yes, to enable layered implementations of OpenGL and D3D.

All of OpenGL, OpenGL ES, and Direct3D 11 require that vertices are written to transform feedback buffers such that flat-shaded attributes have the same value when drawing the contents of the transform feedback buffer as they did in the original drawing when the transform feedback buffer was written (assuming the provoking vertex mode has not changed, in APIs that support more than one mode).

## [](#_version_history)Version History

- Revision 1, (1c) 2021-02-22 (Jesse Hall)
  
  - Added VkPhysicalDeviceProvokingVertexPropertiesEXT::transformFeedbackPreservesTriangleFanProvokingVertex to accommodate implementations that cannot change the transform feedback vertex order for triangle fans.
- Revision 1, (1b) 2020-06-14 (Jesse Hall)
  
  - Added VkPhysicalDeviceProvokingVertexFeaturesEXT::transformFeedbackPreservesProvokingVertex and required that transform feedback write vertices so as to preserve the provoking vertex of each primitive.
- Revision 1, (1a) 2019-10-23 (Jesse Hall)
  
  - Initial draft, based on a proposal by Alexis Hétu

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_provoking_vertex)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0