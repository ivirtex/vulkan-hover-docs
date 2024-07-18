# VK_EXT_provoking_vertex(3) Manual Page

## Name

VK_EXT_provoking_vertex - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

255

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
  target="_blank" rel="noopener">OpenGL / ES support</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jesse Hall <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_provoking_vertex%5D%20@jessehall%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_provoking_vertex%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jessehall</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension allows changing the provoking vertex convention between
Vulkan’s default convention (first vertex) and OpenGL’s convention (last
vertex).

This extension is intended for use by API-translation layers that
implement APIs like OpenGL on top of Vulkan, and need to match the
source API’s provoking vertex convention. Applications using Vulkan
directly should use Vulkan’s default convention.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceProvokingVertexFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProvokingVertexFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceProvokingVertexPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProvokingVertexPropertiesEXT.html)

- Extending
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html):

  - [VkPipelineRasterizationProvokingVertexStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationProvokingVertexStateCreateInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkProvokingVertexModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkProvokingVertexModeEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_PROVOKING_VERTEX_EXTENSION_NAME`

- `VK_EXT_PROVOKING_VERTEX_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROVOKING_VERTEX_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROVOKING_VERTEX_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_PROVOKING_VERTEX_STATE_CREATE_INFO_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) At what granularity should this state be set?

**RESOLVED**: At pipeline bind, with an optional per-render pass
restriction.

The most natural place to put this state is in the graphics pipeline
object. Some implementations require it to be known when creating the
pipeline, and pipeline state is convenient for implementing OpenGL 3.2’s
glProvokingVertex, which can change the state between draw calls.
However, some implementations can only change it approximately render
pass granularity. To accommodate both, provoking vertex will be pipeline
state, but implementations can require that only one mode is used within
a render pass instance; the render pass’s mode is chosen implicitly when
the first pipeline is bound.

2\) Does the provoking vertex mode affect the order that vertices are
written to transform feedback buffers?

**RESOLVED**: Yes, to enable layered implementations of OpenGL and D3D.

All of OpenGL, OpenGL ES, and Direct3D 11 require that vertices are
written to transform feedback buffers such that flat-shaded attributes
have the same value when drawing the contents of the transform feedback
buffer as they did in the original drawing when the transform feedback
buffer was written (assuming the provoking vertex mode has not changed,
in APIs that support more than one mode).

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, (1c) 2021-02-22 (Jesse Hall)

  - Added
    VkPhysicalDeviceProvokingVertexPropertiesEXT::transformFeedbackPreservesTriangleFanProvokingVertex
    to accommodate implementations that cannot change the transform
    feedback vertex order for triangle fans.

- Revision 1, (1b) 2020-06-14 (Jesse Hall)

  - Added
    VkPhysicalDeviceProvokingVertexFeaturesEXT::transformFeedbackPreservesProvokingVertex
    and required that transform feedback write vertices so as to
    preserve the provoking vertex of each primitive.

- Revision 1, (1a) 2019-10-23 (Jesse Hall)

  - Initial draft, based on a proposal by Alexis Hétu

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_provoking_vertex"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
