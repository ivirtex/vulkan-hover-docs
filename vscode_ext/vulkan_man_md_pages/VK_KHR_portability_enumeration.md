# VK_KHR_portability_enumeration(3) Manual Page

## Name

VK_KHR_portability_enumeration - instance extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

395

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_contact" class="anchor"></a>Contact

- Charles Giessen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_portability_enumeration%5D%20@charles-lunarg%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_portability_enumeration%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>charles-lunarg</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-06-02

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- Interacts with
  [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)

**Contributors**  
- Lenny Komow, LunarG

- Charles Giessen, LunarG

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to control whether devices that
expose the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
extension are included in the results of physical device enumeration.
Since devices which support the
[`VK_KHR_portability_subset`](VK_KHR_portability_subset.html) extension
are not fully conformant Vulkan implementations, the Vulkan loader does
not report those devices unless the application explicitly asks for
them. This prevents applications which may not be aware of
non-conformant devices from accidentally using them, as any device which
supports the
[`VK_KHR_portability_subset`](VK_KHR_portability_subset.html) extension
mandates that the extension must be enabled if that device is used.

This extension is implemented in the loader.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_PORTABILITY_ENUMERATION_EXTENSION_NAME`

- `VK_KHR_PORTABILITY_ENUMERATION_SPEC_VERSION`

- Extending [VkInstanceCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateFlagBits.html):

  - `VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-06-02 (Lenny Komow)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_portability_enumeration"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
