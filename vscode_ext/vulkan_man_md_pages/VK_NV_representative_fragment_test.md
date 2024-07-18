# VK_NV_representative_fragment_test(3) Manual Page

## Name

VK_NV_representative_fragment_test - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

167

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Kedarnath Thangudu <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_representative_fragment_test%5D%20@kthangudu%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_representative_fragment_test%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>kthangudu</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-09-13

**Contributors**  
- Kedarnath Thangudu, NVIDIA

- Christoph Kubisch, NVIDIA

- Pierre Boudier, NVIDIA

- Pat Brown, NVIDIA

- Jeff Bolz, NVIDIA

- Eric Werness, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension provides a new representative fragment test that allows
implementations to reduce the amount of rasterization and fragment
processing work performed for each point, line, or triangle primitive.
For any primitive that produces one or more fragments that pass all
other early fragment tests, the implementation is permitted to choose
one or more “representative” fragments for processing and discard all
other fragments. For draw calls rendering multiple points, lines, or
triangles arranged in lists, strips, or fans, the representative
fragment test is performed independently for each of those primitives.

This extension is useful for applications that use an early render pass
to determine the full set of primitives that would be visible in the
final scene. In this render pass, such applications would set up a
fragment shader that enables early fragment tests and writes to an image
or shader storage buffer to record the ID of the primitive that
generated the fragment. Without this extension, the shader would record
the ID separately for each visible fragment of each primitive. With this
extension, fewer stores will be performed, particularly for large
primitives.

The representative fragment test has no effect if early fragment tests
are not enabled via the fragment shader. The set of fragments discarded
by the representative fragment test is implementation-dependent and may
vary from frame to frame. In some cases, the representative fragment
test may not discard any fragments for a given primitive.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html):

  - [VkPipelineRepresentativeFragmentTestStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRepresentativeFragmentTestStateCreateInfoNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceRepresentativeFragmentTestFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRepresentativeFragmentTestFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_REPRESENTATIVE_FRAGMENT_TEST_EXTENSION_NAME`

- `VK_NV_REPRESENTATIVE_FRAGMENT_TEST_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_REPRESENTATIVE_FRAGMENT_TEST_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PIPELINE_REPRESENTATIVE_FRAGMENT_TEST_STATE_CREATE_INFO_NV`

## <a href="#_issues" class="anchor"></a>Issues

\(1\) Is the representative fragment test guaranteed to have any effect?

**RESOLVED**: No. As specified, we only guarantee that each primitive
with at least one fragment that passes prior tests will have one
fragment passing the representative fragment tests. We do not guarantee
that any particular fragment will fail the test.

In the initial implementation of this extension, the representative
fragment test is treated as an optimization that may be completely
disabled for some pipeline states. This feature was designed for a use
case where the fragment shader records information on individual
primitives using shader storage buffers or storage images, with no
writes to color or depth buffers.

\(2\) Will the set of fragments that pass the representative fragment
test be repeatable if you draw the same scene over and over again?

**RESOLVED**: No. The set of fragments that pass the representative
fragment test is implementation-dependent and may vary due to the timing
of operations performed by the GPU.

\(3\) What happens if you enable the representative fragment test with
writes to color and/or depth render targets enabled?

**RESOLVED**: If writes to the color or depth buffer are enabled, they
will be performed for any fragments that survive the relevant tests. Any
fragments that fail the representative fragment test will not update
color buffers. For the use cases intended for this feature, we do not
expect color or depth writes to be enabled.

\(4\) How do derivatives and automatic texture LOD computations work
with the representative fragment test enabled?

**RESOLVED**: If a fragment shader uses derivative functions or texture
lookups using automatic LOD computation, derivatives will be computed
identically whether or not the representative fragment test is enabled.
For the use cases intended for this feature, we do not expect the use of
derivatives in the fragment shader.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2018-09-13 (pbrown)

  - Add issues.

- Revision 1, 2018-08-22 (Kedarnath Thangudu)

  - Internal Revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_representative_fragment_test"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
