# VK\_NV\_representative\_fragment\_test(3) Manual Page

## Name

VK\_NV\_representative\_fragment\_test - device extension



## [](#_registered_extension_number)Registered Extension Number

167

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Kedarnath Thangudu [\[GitHub\]kthangudu](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_representative_fragment_test%5D%20%40kthangudu%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_representative_fragment_test%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-09-13

**Contributors**

- Kedarnath Thangudu, NVIDIA
- Christoph Kubisch, NVIDIA
- Pierre Boudier, NVIDIA
- Pat Brown, NVIDIA
- Jeff Bolz, NVIDIA
- Eric Werness, NVIDIA

## [](#_description)Description

This extension provides a new representative fragment test that allows implementations to reduce the amount of rasterization and fragment processing work performed for each point, line, or triangle primitive. For any primitive that produces one or more fragments that pass all other early fragment tests, the implementation is permitted to choose one or more “representative” fragments for processing and discard all other fragments. For draw calls rendering multiple points, lines, or triangles arranged in lists, strips, or fans, the representative fragment test is performed independently for each of those primitives.

This extension is useful for applications that use an early render pass to determine the full set of primitives that would be visible in the final scene. In this render pass, such applications would set up a fragment shader that enables early fragment tests and writes to an image or shader storage buffer to record the ID of the primitive that generated the fragment. Without this extension, the shader would record the ID separately for each visible fragment of each primitive. With this extension, fewer stores will be performed, particularly for large primitives.

The representative fragment test has no effect if early fragment tests are not enabled via the fragment shader. The set of fragments discarded by the representative fragment test is implementation-dependent and may vary from frame to frame. In some cases, the representative fragment test may not discard any fragments for a given primitive.

## [](#_new_structures)New Structures

- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html):
  
  - [VkPipelineRepresentativeFragmentTestStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRepresentativeFragmentTestStateCreateInfoNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceRepresentativeFragmentTestFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRepresentativeFragmentTestFeaturesNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_REPRESENTATIVE_FRAGMENT_TEST_EXTENSION_NAME`
- `VK_NV_REPRESENTATIVE_FRAGMENT_TEST_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_REPRESENTATIVE_FRAGMENT_TEST_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PIPELINE_REPRESENTATIVE_FRAGMENT_TEST_STATE_CREATE_INFO_NV`

## [](#_issues)Issues

(1) Is the representative fragment test guaranteed to have any effect?

**RESOLVED**: No. As specified, we only guarantee that each primitive with at least one fragment that passes prior tests will have one fragment passing the representative fragment tests. We do not guarantee that any particular fragment will fail the test.

In the initial implementation of this extension, the representative fragment test is treated as an optimization that may be completely disabled for some pipeline states. This feature was designed for a use case where the fragment shader records information on individual primitives using shader storage buffers or storage images, with no writes to color or depth buffers.

(2) Will the set of fragments that pass the representative fragment test be repeatable if you draw the same scene over and over again?

**RESOLVED**: No. The set of fragments that pass the representative fragment test is implementation-dependent and may vary due to the timing of operations performed by the GPU.

(3) What happens if you enable the representative fragment test with writes to color and/or depth render targets enabled?

**RESOLVED**: If writes to the color or depth buffer are enabled, they will be performed for any fragments that survive the relevant tests. Any fragments that fail the representative fragment test will not update color buffers. For the use cases intended for this feature, we do not expect color or depth writes to be enabled.

(4) How do derivatives and automatic texture LOD computations work with the representative fragment test enabled?

**RESOLVED**: If a fragment shader uses derivative functions or texture lookups using automatic LOD computation, derivatives will be computed identically whether or not the representative fragment test is enabled. For the use cases intended for this feature, we do not expect the use of derivatives in the fragment shader.

## [](#_version_history)Version History

- Revision 2, 2018-09-13 (pbrown)
  
  - Add issues.
- Revision 1, 2018-08-22 (Kedarnath Thangudu)
  
  - Internal Revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_representative_fragment_test)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0